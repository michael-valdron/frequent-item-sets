/*
	CSCI4030U: Big Data Project Part 1
	Apriori
	Author: Michael Valdron
	Date: Feb 23, 2018
*/
package apriori

import (
	"fmt"
	"time"

	"github.com/michael-valdron/frequent-item-sets/pkg/common"
)

func Apriori(fname string, t_hold float32) ([]map[int][]int, []map[int]int) {
	k := 0
	var itemsets []map[int][]int
	var counts []map[int]int

	// Get Items First Pass -------------------------------------------------------------------------------------
	start_time := time.Now()
	c_item_counts, min_supp, _, _ := common.GetFreqItems(fname, t_hold, false)

	itemsets = []map[int][]int{}
	itemsets = append(itemsets, make(map[int][]int))

	common.FilterItems(itemsets, c_item_counts, min_supp, k)

	counts = append(counts, c_item_counts)

	finish_time := time.Now()

	fmt.Printf("Pass %d took: %d minutes and %d seconds.\n",
		(k + 1),
		int(finish_time.Sub(start_time).Minutes()),
		(int(finish_time.Sub(start_time).Seconds()) % 60))

	// ----------------------------------------------------------------------------------------------------------

	// Calculate Itemsets Second to Finite Pass -----------------------------------------------------------------
	for len(itemsets[k]) > 0 && k < 2 {
		start_time = time.Now()
		c_itemsets, c_itemset_counts := common.GenTuples(itemsets, k, map[int]bool{}, 0, false)
		c_itemsets, c_itemset_counts = common.GetFreqTuples(c_itemsets, c_itemset_counts, fname, k, min_supp, map[int]bool{}, 0)

		itemsets = append(itemsets, make(map[int][]int))

		common.FilterItemsets(itemsets, c_itemsets, c_itemset_counts, min_supp, k)

		counts = append(counts, c_itemset_counts)
		finish_time = time.Now()
		k++
		fmt.Printf("Pass %d took: %d minutes and %d seconds.\n",
			(k + 1),
			int(finish_time.Sub(start_time).Minutes()),
			(int(finish_time.Sub(start_time).Seconds()) % 60))
	}
	if len(itemsets[k]) == 0 {
		itemsets = itemsets[0:k]
		counts = counts[0:k]
	}
	// ----------------------------------------------------------------------------------------------------------

	return itemsets, counts
}
