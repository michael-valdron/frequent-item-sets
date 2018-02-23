/*
	CSCI4030U: Big Data Project Part 1
	Apriori
	Author: Michael Valdron
	Date: Feb 23, 2018
*/
package main

func Apriori(fname string, t_hold float32) ([]map[int][]int, []map[int]int) {
	k := 0
	var itemsets []map[int][]int
	var counts []map[int]int

	// Get Items First Pass -------------------------------------------------------------------------------------
	c_item_counts, min_supp, _, _ := GetFreqItems(fname, t_hold, false)

	itemsets = []map[int][]int{}
	itemsets = append(itemsets, make(map[int][]int))

	FilterItems(itemsets, c_item_counts, min_supp, k)

	counts = append(counts, c_item_counts)

	// ----------------------------------------------------------------------------------------------------------

	// Calculate Itemsets Second to Finite Pass -----------------------------------------------------------------
	for len(itemsets[k]) > 0 && k < 2 {
		c_itemsets, c_itemset_counts := GenTuples(itemsets, k, map[int]bool{}, 0, false)
		c_itemsets, c_itemset_counts = GetFreqTuples(c_itemsets, c_itemset_counts, fname, k, min_supp, map[int]bool{}, 0)

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
