/*
	CSCI4030U: Big Data Project Part 1
	PCY
	Author: Michael Valdron
	Date: Feb 12, 2018
*/
package main

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func mulHashFunc(value string, m int) int {
	var a float64 = (math.Sqrt(5) - 1) / 2
	var h float64 = 0.0

	for _, s := range strings.Split(value, ",") {
		n, _ := strconv.Atoi(s)
		h += float64(n)
	}

	prod := h * a
	frac := prod - math.Floor(prod)

	return int(math.Floor(float64(m) * frac))
}

func hashBasket(basket []string, itemset_hashtable []int, n_bins int) {
	for i := range basket {
		for j := i; j < len(basket); j++ {
			if basket[i] != basket[j] {
				pair := basket[i] + "," + basket[j]
				itemset_hashtable[mulHashFunc(pair, n_bins)] += 1
			}
		}
	}
}

func hashtableToBitmap(itemset_hashtable []int, itemset_bitmap []bool, min_supp int) {
	for i := range itemset_hashtable {
		itemset_bitmap[i] = itemset_hashtable[i] >= min_supp
	}
}

func getBasketsPCY(k_items []string, init_items []string, bitmap []bool, n_bins int) []string {
	var baskets []string
	for _, k_item_set_str := range k_items {
		k_item_set := strings.Split(k_item_set_str, TUPLE_SEP)
		for _, item := range init_items {
			if !checkIn(item, k_item_set) {
				sel_item_set := append(k_item_set, item)
				sort.Strings(sel_item_set)
				sel_item_set_str := strings.Join(sel_item_set, TUPLE_SEP)
				if bitmap[mulHashFunc(sel_item_set_str, n_bins)] && !checkIn(sel_item_set_str, baskets) {
					baskets = append(baskets, sel_item_set_str)
				}
			}
		}
	}

	return baskets
}

func PCY(fname string, t_hold float32) []map[string]int {
	k := 0
	n := float32(getFileLength(fname))
	min_supp := int(t_hold * n)
	bins := int(n * 0.2)
	var itemset_counts []map[string]int

	items, itemset_bitmap := getFreqItems(fname, min_supp, true, bins)
	itemset_counts = append(itemset_counts, items)
	items = nil

	// Debug message
	//fmt.Println(itemset_bitmap)

	for len(itemset_counts[k]) > 0 && k < 2 {
		k_item_sets := getKeyStr(itemset_counts[k])
		init_item_sets := getKeyStr(itemset_counts[0])
		candidate_item_set := getBasketsPCY(k_item_sets, init_item_sets, itemset_bitmap, bins)
		itemset_counts = append(itemset_counts, getFreqTuples(fname, candidate_item_set, min_supp))
		k += 1
	}

	return itemset_counts
}
