/*
	CSCI4030U: Big Data Project Part 1
	Common Functions
	Author: Michael Valdron
	Date: Feb 12, 2018
*/
package main

import (
	"bufio"
	"fmt"
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
	for _, v := range list {
		if i == v {
			return true
		}
	}
	return false
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
func GetFreqItems(fname string, t_hold float32, is_pcy bool) (map[int]int, int, map[int]bool) {
	var n float32
	f, err := os.Open(fname)
	defer f.Close()
	if err != nil {
		fmt.Printf("%s\n", err)
		return map[int]int{}, 0, map[int]bool{}
	}

	n = 0
	fs := bufio.NewScanner(f)
	c_items := make(map[int]bool)
	c_item_counts := make(map[int]int)
	itemsets_hashtable := make([]int, N_BINS)
	itemsets_bitmap := make(map[int]bool)

	for fs.Scan() {
		str_line := strings.Split(strings.TrimSuffix(fs.Text(), FILE_SEP), FILE_SEP)
		GetItemCounts(str_line, c_items, c_item_counts)
		if is_pcy {
			HashBasket(str_line, itemsets_hashtable, N_BINS)
		}
		n++
	}

	// Calculate minimum support
	min_supp := int(t_hold * n)

	if is_pcy {
		itemsets_bitmap = HashtableToBitmap(itemsets_hashtable, min_supp)
	}

	return c_item_counts, min_supp, itemsets_bitmap
}

func GetItemCounts(tuple []string, c_items map[int]bool, c_item_counts map[int]int) {
	for _, item := range tuple {
		int_item, _ := strconv.Atoi(item)
		if !c_items[int_item] {
			c_items[int_item] = true
			c_item_counts[int_item] = 1
		} else {
			c_item_counts[int_item] += 1
		}
	}
}

func GenTuples(itemsets []map[int][]int, k int, bitmap map[int]bool, is_pcy bool) (map[int][]int, map[int]int) {
	// Get Frequent Tuples
	c_itemsets := make(map[int][]int)
	c_itemset_counts := make(map[int]int)

	// Generate Tuples
	row_id := 1
	for _, k_itemset := range itemsets[k] {
		for item := range itemsets[0] {
			if is_pcy {
				AddItemToSetPCY(k_itemset, item, c_itemsets, c_itemset_counts, &row_id, k, bitmap)
			} else {
				AddItemToSetApriori(k_itemset, item, c_itemsets, c_itemset_counts, &row_id, k)
			}
		}
	}

	return c_itemsets, c_itemset_counts
}

func AddItemToSetApriori(selSet []int, item int, c_itemsets map[int][]int, c_itemset_counts map[int]int, row_id *int, k int) {
	if !CheckIn(item, selSet) {
		itemset := append(selSet, item)
		is_similar := false
		for _, c_itemset := range c_itemsets {
			is_similar = IsSimilar(itemset, c_itemset, k)
			if is_similar {
				break
			}
		}
		if !is_similar {
			c_itemsets[*row_id] = itemset
			c_itemset_counts[*row_id] = 0
			*row_id++
		}
	}
}

func AddItemToSetPCY(selSet []int, item int, c_itemsets map[int][]int, c_itemset_counts map[int]int, row_id *int, k int, bitmap map[int]bool) {
	if !CheckIn(item, selSet) {
		itemset := append(selSet, item)
		if bitmap[MulHashFunc(itemset, N_BINS)] {
			is_similar := false
			for _, c_itemset := range c_itemsets {
				is_similar = IsSimilar(itemset, c_itemset, k)
				if is_similar {
					break
				}
			}
			if !is_similar {
				c_itemsets[*row_id] = itemset
				c_itemset_counts[*row_id] = 0
				*row_id++
			}
		}
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

func GetFreqTuples(itemsets map[int][]int, itemset_counts map[int]int, fname string, k int) (map[int][]int, map[int]int) {
	var c_itemsets map[int][]int
	var c_itemset_counts map[int]int
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

	fs := bufio.NewScanner(f)

	for fs.Scan() {
		line_map := GetLineBitmap(strings.Split(strings.TrimSuffix(fs.Text(), FILE_SEP), FILE_SEP))
		GetItemsetCounts(c_itemsets, c_itemset_counts, line_map, k)
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

func GetItemsetCounts(c_itemsets map[int][]int, c_itemset_counts map[int]int, line_map map[int]bool, k int) {
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
