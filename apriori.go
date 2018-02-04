/*
	CSCI4030U: Big Data Project Part 1
	Apriori
	Author: Michael Valdron
	Date: Feb 12, 2018
*/
package main

import (
	"sort"
	"strings"
)

func getFreqItemsApriori(fname string, min_supp int) map[string]int {
	itemset, _ := getFreqItems(fname, min_supp, false, 0)
	return itemset
}

func getBasketsApriori(k_items []string, init_items []string) []string {
	var baskets []string
	for _, k_item_set_str := range k_items {
		k_item_set := strings.Split(k_item_set_str, TUPLE_SEP)
		for _, item := range init_items {
			if !checkIn(item, k_item_set) {
				sel_item_set := append(k_item_set, item)
				sort.Strings(sel_item_set)
				sel_item_set_str := strings.Join(sel_item_set, TUPLE_SEP)
				if !checkIn(sel_item_set_str, baskets) {
					baskets = append(baskets, sel_item_set_str)
				}
			}
		}
	}

	return baskets
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
		candidate_item_set := getBasketsApriori(k_item_sets, init_item_sets)
		itemset_counts = append(itemset_counts, getFreqTuples(fname, candidate_item_set, min_supp))

		k += 1
	}

	return itemset_counts
}
