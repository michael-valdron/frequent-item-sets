/*
	CSCI4030U: Big Data Project Part 1
	Apriori
	Author: Michael Valdron
	Date: Feb 12, 2018
*/
package main

func Apriori(fname string, t_hold float32) ([]map[int][]int, []map[int]int) {
	k := 0
	var n float32
	var itemsets []map[int][]int
	var counts []map[int]int

	// Get Items First Pass -------------------------------------------------------------------------------------

	c_item_counts, n := GetFreqItems(fname)

	// Calculate minimum support
	min_supp := int(t_hold * n)

	itemsets = []map[int][]int{}
	itemsets = append(itemsets, make(map[int][]int))

	FilterItems(itemsets, c_item_counts, min_supp, k)

	counts = append(counts, c_item_counts)

	// ----------------------------------------------------------------------------------------------------------

	// Calculate Itemsets Second to Finite Pass -----------------------------------------------------------------
	for len(itemsets[k]) > 0 && k < 2 {
		c_itemsets, c_itemset_counts := GenTuples(itemsets, k)
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
