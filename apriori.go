package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const FILE_SEP = " "

func CheckIn(i int, list []int) bool {
	for _, v := range list {
		if i == v {
			return true
		}
	}
	return false
}

func Apriori(fname string, t_hold float32) ([]map[int][]int, []map[int]int) {
	k := 0
	var n float32
	var itemsets []map[int][]int
	var counts []map[int]int

	// Get Items First Pass -------------------------------------------------------------------------------------
	f, err := os.Open(fname)
	defer f.Close()
	if err != nil {
		fmt.Printf("%s\n", err)
		return []map[int][]int{}, []map[int]int{}
	}

	n = 0
	fs := bufio.NewScanner(f)
	c_items := make(map[int]bool)
	c_item_counts := make(map[int]int)

	for fs.Scan() {
		for _, item := range strings.Split(strings.TrimSuffix(fs.Text(), FILE_SEP), FILE_SEP) {
			int_item, _ := strconv.Atoi(item)
			if !c_items[int_item] {
				c_items[int_item] = true
				c_item_counts[int_item] = 1
			} else {
				c_item_counts[int_item] += 1
			}
		}
		n++
	}

	// Calculate minimum support
	min_supp := int(t_hold * n)

	itemsets = []map[int][]int{}
	itemsets = append(itemsets, make(map[int][]int))
	// Filter items
	for item, count := range c_item_counts {
		if count < min_supp {
			delete(c_item_counts, item)
		} else {
			itemsets[k][item] = []int{item}
		}
	}

	counts = append(counts, c_item_counts)

	// ----------------------------------------------------------------------------------------------------------

	for len(itemsets[k]) > 0 && k < 2 {
		// Get Frequent Tuples
		c_itemsets := make(map[int][]int)
		c_itemset_counts := make(map[int]int)

		// Generate Tuples
		i := 1
		for _, k_itemset := range itemsets[k] {
			for item := range itemsets[0] {
				if !CheckIn(item, k_itemset) {
					itemset := append(k_itemset, item)
					is_similar := false
					for _, c_itemset := range c_itemsets {
						n_items_similar := 0
						for _, v := range itemset {
							if CheckIn(v, c_itemset) {
								n_items_similar++
							}
						}
						if n_items_similar > (k + 1) {
							is_similar = true
							break
						}
					}
					if !is_similar {
						c_itemsets[i] = itemset
						c_itemset_counts[i] = 0
						i++
					}
				}
			}
		}

		// Count Tuples
		f, err := os.Open(fname)
		defer f.Close()
		if err != nil {
			fmt.Printf("%s\n", err)
			return []map[int][]int{}, []map[int]int{}
		}

		fs := bufio.NewScanner(f)

		for fs.Scan() {
			line_map := make(map[int]bool)
			for _, item := range strings.Split(strings.TrimSuffix(fs.Text(), FILE_SEP), FILE_SEP) {
				int_item, _ := strconv.Atoi(item)
				line_map[int_item] = true
			}
			for key, tuple := range c_itemsets {
				n_similar := 0
				for _, item := range tuple {
					if line_map[item] {
						n_similar++
					} else {
						break
					}
				}
				if n_similar > (k + 1) {
					c_itemset_counts[key]++
				}
			}
		}

		itemsets = append(itemsets, make(map[int][]int))
		// Filter items
		for key, count := range c_itemset_counts {
			if count < min_supp {
				delete(c_itemset_counts, key)
				delete(c_itemsets, key)
			} else {
				itemsets[k+1][key] = c_itemsets[key]
			}
		}

		counts = append(counts, c_itemset_counts)

		k++
	}
	if len(itemsets[k]) == 0 {
		itemsets = itemsets[0:k]
		counts = counts[0:k]
	}

	return itemsets, counts
}
