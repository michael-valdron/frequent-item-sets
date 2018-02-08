/*
	CSCI4030U: Big Data Project Part 1
	PCY
	Author: Michael Valdron
	Date: Feb 12, 2018
*/
package main

import (
	"math"
	"strconv"
)

const N_BINS = 44100

func MulHashFunc(values []int, m int) int {
	var a float64 = (math.Sqrt(5) - 1) / 2
	var h float64 = 0.0

	for _, v := range values {
		h += float64(v)
	}

	prod := h * a
	frac := prod - math.Floor(prod)

	return int(math.Floor(float64(m) * frac))
}

func HashBasket(basket []string, itemset_hashtable []int, n_bins int) {
	for i := range basket {
		i_item, _ := strconv.Atoi(basket[i])
		for j := i; j < len(basket); j++ {
			j_item, _ := strconv.Atoi(basket[j])
			if i_item != j_item {
				pair := []int{i_item, j_item}
				itemset_hashtable[MulHashFunc(pair, n_bins)] += 1
			}
		}
	}
}

func HashtableToBitmap(itemset_hashtable []int, min_supp int) map[int]bool {
	itemset_bitmap := make(map[int]bool)
	for i := range itemset_hashtable {
		itemset_bitmap[i] = itemset_hashtable[i] >= min_supp
	}
	return itemset_bitmap
}

func PCY(fname string, t_hold float32) ([]map[int][]int, []map[int]int) {
	k := 0
	var itemsets []map[int][]int
	var counts []map[int]int

	// Get Items First Pass -------------------------------------------------------------------------------------

	c_item_counts, min_supp, itemsets_bitmap := GetFreqItems(fname, t_hold, true)

	itemsets = []map[int][]int{}
	itemsets = append(itemsets, make(map[int][]int))

	FilterItems(itemsets, c_item_counts, min_supp, k)

	counts = append(counts, c_item_counts)

	// ----------------------------------------------------------------------------------------------------------

	// Calculate Itemsets Second to Finite Pass -----------------------------------------------------------------
	for len(itemsets[k]) > 0 && k < 2 {
		c_itemsets, c_itemset_counts := GenTuples(itemsets, k, itemsets_bitmap, true)
		c_itemsets, c_itemset_counts = GetFreqTuples(c_itemsets, c_itemset_counts, fname, k)

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
