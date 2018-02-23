/*
	CSCI4030U: Big Data Project Part 1
	Common Functions
	Author: Michael Valdron
	Date: Feb 23, 2018
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const FILE_SEP = " "

func Min(list []int) int {
	min := 0
	for _, v := range list {
		if v > min {
			min = v
		}
	}
	return min
}

func Max(list []int) int {
	max := 0
	for _, v := range list {
		if v < max {
			max = v
		}
	}
	return max
}

func CheckIn(i int, list []int) bool {
	for k := range list {
		if i == list[k] {
			return true
		}
	}
	return false
}

func GetLineCount(fname string) int {
	f, err := os.Open(fname)
	defer f.Close()
	if err != nil {
		fmt.Printf("%s\n", err)
		return 0
	}

	lc := 0
	fr := bufio.NewReader(f)
	_, _, err = fr.ReadLine()

	// Get line count of file
	for err != io.EOF {
		lc++
		_, _, err = fr.ReadLine()
	}

	return lc
}

func FilterItems(itemsets []map[int][]int, c_item_counts map[int]int, min_supp int, k int) {
	// Filter items
	for item, count := range c_item_counts {
		if count < min_supp {
			delete(c_item_counts, item)
		} else {
			itemsets[k][item] = []int{item}
		}
	}
}

func FilterItemsets(itemsets []map[int][]int, c_itemsets map[int][]int, c_itemset_counts map[int]int, min_supp int, k int) {
	// Filter itemsets
	for key, count := range c_itemset_counts {
		if count < min_supp {
			delete(c_itemset_counts, key)
			delete(c_itemsets, key)
		} else {
			itemsets[k+1][key] = c_itemsets[key]
		}
	}
}

/*
   GetFreqItems
   ============

   Gets the frequent items and the counts
   for the first pass.  Keys returned are
   the items and the values are the counts.
   Also returns the number of baskets in the
   dataset.
*/
func GetFreqItems(fname string, t_hold float32, is_pcy bool) (map[int]int, int, int, map[int]bool) {
	n := GetLineCount(fname)
	if n == 0 {
		return map[int]int{}, 0, 0, map[int]bool{}
	}
	f, err := os.Open(fname)
	defer f.Close()
	if err != nil {
		fmt.Printf("%s\n", err)
		return map[int]int{}, 0, 0, map[int]bool{}
	}

	n_bins := int(float64(n) * float64(0.3))
	c_items := make(map[int]bool)
	c_item_counts := make(map[int]int)
	itemsets_hashtable := make([]int, n_bins)
	itemsets_bitmap := make(map[int]bool)
	fr := bufio.NewReader(f)
	line, _, err := fr.ReadLine()
	for err != io.EOF {
		int_line := []int{}
		for _, str_item := range strings.Split(strings.TrimSuffix(string(line), FILE_SEP), FILE_SEP) {
			int_item, _ := strconv.Atoi(str_item)
			int_line = append(int_line, int_item)
		}
		GetItemCounts(int_line, c_items, c_item_counts, itemsets_hashtable, n_bins, is_pcy)
		line, _, err = fr.ReadLine()
	}

	// Calculate minimum support
	min_supp := int(t_hold * float32(n))

	if is_pcy {
		HashtableToBitmap(itemsets_hashtable, itemsets_bitmap, min_supp)
	}

	return c_item_counts, min_supp, n_bins, itemsets_bitmap
}

func GetItemCounts(tuple []int,
	c_items map[int]bool,
	c_item_counts map[int]int,
	itemset_hashtable []int,
	n_bins int,
	is_pcy bool) {
	n := len(tuple)
	for i, i_item := range tuple {
		if !c_items[i_item] {
			c_items[i_item] = true
			c_item_counts[i_item] = 1
		} else {
			c_item_counts[i_item] += 1
		}
		if is_pcy {
			HashPair(tuple, itemset_hashtable, i, i_item, n, n_bins)
		}
	}
}

func GenTuples(itemsets []map[int][]int,
	k int,
	bitmap map[int]bool,
	n_bins int,
	is_pcy bool) (map[int][]int, map[int]int) {
	// Get Frequent Tuples
	items := make(map[int]bool)
	c_itemsets := make(map[int][]int)
	c_itemset_counts := make(map[int]int)

	for _, k_itemset := range itemsets[k] {
		for _, item := range k_itemset {
			if !items[item] {
				items[item] = true
			}
		}
	}

	// Generate Tuples
	row_id := 1
	for _, a := range itemsets[k] {
		selSet := a
		for b := range items {
			if !CheckIn(b, selSet) {
				selSet = append(selSet, b)
				if len(selSet) > (k + 1) {
					if is_pcy && k < 1 {
						if bitmap[MulHashFunc(selSet, n_bins)] {
							AddItemToSet(selSet, c_itemsets, c_itemset_counts, &row_id, k)
						}
					} else {
						AddItemToSet(selSet, c_itemsets, c_itemset_counts, &row_id, k)
					}
					selSet = a
				}
			}
		}
	}

	return c_itemsets, c_itemset_counts
}

func AddItemToSet(selSet []int, c_itemsets map[int][]int, c_itemset_counts map[int]int, row_id *int, k int) {
	is_similar := false
	for _, c_itemset := range c_itemsets {
		is_similar = IsSimilar(selSet, c_itemset, k)
		if is_similar {
			break
		}
	}
	if !is_similar {
		c_itemsets[*row_id] = selSet
		c_itemset_counts[*row_id] = 0
		*row_id++
	}
}

func IsSimilar(itemset []int, c_itemset []int, k int) bool {
	n_items_similar := 0
	for _, v := range itemset {
		if CheckIn(v, c_itemset) {
			n_items_similar++
		}
	}
	return n_items_similar > (k + 1)
}

func GetFreqTuples(itemsets map[int][]int,
	itemset_counts map[int]int,
	fname string,
	k int,
	min_supp int,
	itemsets_bitmap map[int]bool,
	n_bins int) (map[int][]int, map[int]int) {
	var c_itemsets map[int][]int
	var c_itemset_counts map[int]int
	var itemsets_hashtable []int
	// Count Tuples
	f, err := os.Open(fname)
	defer f.Close()
	if err != nil {
		fmt.Printf("%s\n", err)
		return map[int][]int{}, map[int]int{}
	}

	c_itemsets = make(map[int][]int)
	c_itemset_counts = make(map[int]int)

	for key := range itemsets {
		c_itemsets[key] = itemsets[key]
	}

	for key := range itemset_counts {
		c_itemset_counts[key] = itemset_counts[key]
	}

	itemsets_hashtable = make([]int, n_bins)
	fr := bufio.NewReader(f)
	line, _, err := fr.ReadLine()

	for err != io.EOF {
		f_line := strings.Split(strings.TrimSuffix(string(line), FILE_SEP), FILE_SEP)
		line_map := GetLineBitmap(f_line)
		GetItemsetCounts(f_line, c_itemsets, c_itemset_counts, line_map, k, itemsets_hashtable, n_bins)
		line, _, err = fr.ReadLine()
	}

	return c_itemsets, c_itemset_counts
}

func GetLineBitmap(line []string) map[int]bool {
	line_map := make(map[int]bool)
	for _, item := range line {
		int_item, _ := strconv.Atoi(item)
		line_map[int_item] = true
	}
	return line_map
}

func GetItemsetCounts(f_line []string,
	c_itemsets map[int][]int,
	c_itemset_counts map[int]int,
	line_map map[int]bool,
	k int,
	itemset_hashtable []int,
	n_bins int) {
	for key, tuple := range c_itemsets {
		if IsSubset(tuple, line_map, k) {
			c_itemset_counts[key]++
		}
	}
}

func IsSubset(tuple []int, line_map map[int]bool, k int) bool {
	n_similar := 0
	for _, item := range tuple {
		if line_map[item] {
			n_similar++
		} else {
			break
		}
	}
	return n_similar > (k + 1)
}
