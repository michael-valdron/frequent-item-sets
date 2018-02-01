/*
	CSCI4030U: Big Data Project Part 1
	Apriori
	Author: Michael Valdron
	Date: Feb 12, 2018
*/
package main

import (
	"bufio"
	"sort"
	"strings"
)

func getFreqItemsApriori(fname string, min_supp int) map[string]int {
	f, is_open := openFile(fname)

	if is_open {
		defer f.Close()
		item_counts := make(map[string]int)
		fs := bufio.NewScanner(f)
		keys := []string{}

		for fs.Scan() {
			basket := strings.Split(strings.TrimSuffix(fs.Text(), " "), " ")
			for _, item := range basket {
				if len(keys) == 0 || !checkIn(item, keys) {
					keys = append(keys, item)
					item_counts[item] = 1
				} else {
					item_counts[item] += 1
				}
			}
		}

		filterItemsets(item_counts, min_supp)

		return item_counts
	} else {
		return map[string]int{}
	}
}

func getBasketsApriori(k_items []string, init_items []string, k int) []string {
	var baskets []string
	if k == 2 {
		for i, k_item_set := range k_items {
			for j := i; j < len(init_items); j++ {
				if !checkIn(init_items[j], strings.Split(k_item_set, ",")) {
					sel_item_set := append(strings.Split(k_item_set, ","), init_items[j])
					sort.Strings(sel_item_set)
					if !checkIn(strings.Join(sel_item_set, ","), baskets) {
						baskets = append(baskets, strings.Join(sel_item_set, ","))
					}
				}
			}
		}
	} else {
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
	}

	return baskets
}

func getFreqTuplesApriori(fname string, tuples []string, min_supp int) map[string]int {
	f, is_open := openFile(fname)

	if is_open {
		defer f.Close()
		itemset_counts := make(map[string]int)
		keys := []string{}

		for _, tuple := range tuples {
			f.Seek(0, 0)
			fs := bufio.NewScanner(f)
			items := strings.Split(tuple, ",")

			for fs.Scan() {
				basket := strings.Split(strings.TrimSuffix(fs.Text(), " "), " ")
				similar_count := 0
				for _, item := range items {
					if checkIn(item, basket) {
						similar_count += 1
					}
				}
				if similar_count == len(items) {
					if !checkIn(tuple, keys) {
						keys = append(keys, tuple)
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

func Apriori(fname string, t_hold float32) []map[string]int {
	k := 0
	n := float32(getFileLength(fname))
	min_supp := int(t_hold * n)
	var itemset_counts []map[string]int

	itemset_counts = append(itemset_counts, getFreqItemsApriori(fname, min_supp))

	for len(itemset_counts[k]) > 0 && k < 2 {
		k_item_sets := getKeyStr(itemset_counts[k])
		init_item_sets := getKeyStr(itemset_counts[0])
		candidate_item_set := getBasketsApriori(k_item_sets, init_item_sets, k)
		itemset_counts = append(itemset_counts, getFreqTuplesApriori(fname, candidate_item_set, min_supp))
		k += 1
	}

	return itemset_counts
}
