/*
	CSCI4030U: Big Data Project Part 1
	PCY
	Author: Michael Valdron
	Date: Feb 12, 2018
*/
package main

import (
	"bufio"
	"sort"
	"strconv"
	"strings"
)

const N_BINS = 500

func hashFuncA(value string, m int) int {
	h := 0

	for _, s := range strings.Split(value, ",") {
		n, _ := strconv.Atoi(s)
		h += n
	}

	return h % m
}

func hashtableToBitmap(itemset_hashtable []int, itemset_bitmap []bool, min_supp int) {
	for i := range itemset_hashtable {
		itemset_bitmap[i] = itemset_hashtable[i] >= min_supp
	}
}

func getFreqItemsPCY(fname string, min_supp int, itemset_bitmap []bool) map[string]int {
	f, is_open := openFile(fname)

	if is_open {
		defer f.Close()
		item_counts := make(map[string]int)
		itemset_hashtable := make([]int, len(itemset_bitmap))
		fs := bufio.NewScanner(f)

		for fs.Scan() {
			bucket := strings.Split(strings.TrimSuffix(fs.Text(), " "), " ")
			for _, item := range bucket {
				keys := getKeyStr(item_counts)
				if len(keys) == 0 || !checkIn(item, keys) {
					item_counts[item] = 1
				} else {
					item_counts[item] += 1
				}
			}
			for i := range bucket {
				for j := i; j < len(bucket); j++ {
					if bucket[i] != bucket[j] {
						pair := bucket[i] + "," + bucket[j]
						itemset_hashtable[hashFuncA(pair, len(itemset_bitmap))] += 1
					}
				}
			}
		}

		// Debug message
		//fmt.Println(itemset_hashtable)

		filterItemsets(item_counts, min_supp)
		hashtableToBitmap(itemset_hashtable, itemset_bitmap, min_supp)

		return item_counts
	} else {
		return map[string]int{}
	}
}

func getBasketsPCY(k_items []string, init_items []string) []string {
	var baskets []string
	for _, k_item_set := range k_items {
		for _, item := range init_items {
			if !checkIn(item, strings.Split(k_item_set, ",")) {
				sel_item_set := append(strings.Split(k_item_set, ","), item)
				sort.Strings(sel_item_set)
				if !checkIn(strings.Join(sel_item_set, ","), baskets) {
					baskets = append(baskets, strings.Join(sel_item_set, ","))
				}
			}
		}
	}

	return baskets
}

func getFreqTuplesPCY(fname string, tuples []string, min_supp int) map[string]int {
	f, is_open := openFile(fname)

	if is_open {
		defer f.Close()
		itemset_counts := make(map[string]int)

		for _, tuple := range tuples {
			f.Seek(0, 0)
			fs := bufio.NewScanner(f)
			for fs.Scan() {
				similar_count := 0
				for _, item := range strings.Split(tuple, ",") {
					if checkIn(item, strings.Split(strings.TrimSuffix(fs.Text(), " "), " ")) {
						similar_count += 1
					}
				}
				if similar_count == len(strings.Split(tuple, ",")) {
					keys := getKeyStr(itemset_counts)
					if !checkIn(tuple, keys) {
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

func PCY(fname string, t_hold int) []map[string]int {
	k := 0
	min_supp := t_hold
	itemset_bitmap := make([]bool, N_BINS)
	var itemset_counts []map[string]int

	itemset_counts = append(itemset_counts, getFreqItemsPCY(fname, min_supp, itemset_bitmap))

	// Debug message
	//fmt.Println(itemset_bitmap)

	for len(itemset_counts[k]) > 0 && k < 2 {
		k_item_sets := getKeyStr(itemset_counts[k])
		init_item_sets := getKeyStr(itemset_counts[0])
		candidate_item_set := getBasketsPCY(k_item_sets, init_item_sets)
		itemset_counts = append(itemset_counts, getFreqTuplesPCY(fname, candidate_item_set, min_supp))
		k += 1
	}

	return itemset_counts
}
