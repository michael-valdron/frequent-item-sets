/*
	CSCI4030U: Big Data Project Part 1
	Apriori and PCY Helper Functions
	Author: Michael Valdron
	Date: Feb 12, 2018
*/

package main

import (
	"bufio"
	"strings"
)

const TUPLE_SEP = ","
const FILE_TUPLE_SEP = " "

// Apriori and PCY algorithm helper functions

func filterItemsets(itemset_counts map[string]int, min_supp int) {
	for key, value := range itemset_counts {
		if value < min_supp {
			delete(itemset_counts, key)
		}
	}
}

func getFreqItems(fname string, min_supp int, is_pcy bool, n_bins int) (map[string]int, []bool) {
	f, is_open := openFile(fname)

	if is_open {
		defer f.Close()
		item_counts := make(map[string]int)
		fs := bufio.NewScanner(f)
		keys := set{make(map[string]bool)}
		itemset_hashtable := make([]int, n_bins)
		itemset_bitmap := make([]bool, n_bins)

		for fs.Scan() {
			basket := strings.Split(strings.TrimSuffix(fs.Text(), FILE_TUPLE_SEP), FILE_TUPLE_SEP)
			for _, item := range basket {
				if keys.Add(item) {
					item_counts[item] = 1
				} else {
					item_counts[item] += 1
				}
			}

			if is_pcy {
				hashBasket(basket, itemset_hashtable, n_bins)
			}
		}

		filterItemsets(item_counts, min_supp)
		if is_pcy {
			hashtableToBitmap(itemset_hashtable, itemset_bitmap, min_supp)
		}

		return item_counts, itemset_bitmap
	} else {
		return map[string]int{}, []bool{}
	}
}

func getFreqTuples(fname string, tuples []string, min_supp int) map[string]int {
	f, is_open := openFile(fname)

	if is_open {
		defer f.Close()
		itemset_counts := make(map[string]int)
		keys := set{make(map[string]bool)}
		fs := bufio.NewScanner(f)

		for fs.Scan() {
			basket := SplitToSet(fs.Text(), FILE_TUPLE_SEP)
			for _, tuple := range tuples {
				items := strings.Split(tuple, TUPLE_SEP)
				if basket.IsSuperset(items) {
					if keys.Add(tuple) {
						itemset_counts[tuple] = 1
					} else {
						itemset_counts[tuple] += 1
					}
				}
			}
		}

		filterItemsets(itemset_counts, min_supp)

		return itemset_counts
	} else {
		return map[string]int{}
	}
}
