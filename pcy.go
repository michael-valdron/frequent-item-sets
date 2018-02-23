/*
	CSCI4030U: Big Data Project Part 1
	PCY
	Author: Michael Valdron
	Date: Feb 23, 2018
*/
package main

import (
	"math"
)

func MulHashFunc(values []int, m int) int {
	const a = 0.618033989
	var h float64 = 0.0

	for i := 0; i < len(values); i++ {
		h += float64(values[i])
	}

	prod := h * a
	frac := prod - math.Floor(prod)

	return int(math.Floor(float64(m) * frac))
}

func HashPair(basket []int, itemset_hashtable []int, index int, i_item int, n int, n_bins int) {
	for j := index; j < n; j++ {
		j_item := basket[j]
		if i_item != j_item {
			pair := []int{i_item, j_item}
			itemset_hashtable[MulHashFunc(pair, n_bins)] += 1
		}
	}
}

func HashtableToBitmap(itemset_hashtable []int, itemset_bitmap map[int]bool, min_supp int) {
	for i := range itemset_hashtable {
		itemset_bitmap[i] = itemset_hashtable[i] >= min_supp
	}
}

func PCY(fname string, t_hold float32) ([]map[int][]int, []map[int]int) {
	k := 0
	var itemsets []map[int][]int
	var counts []map[int]int

	// Get Items First Pass -------------------------------------------------------------------------------------
	c_item_counts, min_supp, n_bins, itemsets_bitmap := GetFreqItems(fname, t_hold, true)
	itemsets = []map[int][]int{}
	itemsets = append(itemsets, make(map[int][]int))

	FilterItems(itemsets, c_item_counts, min_supp, k)

	counts = append(counts, c_item_counts)

	// ----------------------------------------------------------------------------------------------------------

	// Calculate Itemsets Second to Finite Pass -----------------------------------------------------------------
	for len(itemsets[k]) > 0 && k < 2 {
		c_itemsets, c_itemset_counts := GenTuples(itemsets, k, itemsets_bitmap, n_bins, true)
		c_itemsets, c_itemset_counts = GetFreqTuples(c_itemsets, c_itemset_counts, fname, k, min_supp, itemsets_bitmap, n_bins)

		itemsets = append(itemsets, make(map[int][]int))

		FilterItemsets(itemsets, c_itemsets, c_itemset_counts, min_supp, k)

		counts = append(counts, c_itemset_counts)

		k++
	}
	if len(itemsets[k]) == 0 {
		itemsets = itemsets[0:k]
		counts = counts[0:k]
	}
	// ----------------------------------------------------------------------------------------------------------

	return itemsets, counts
}
