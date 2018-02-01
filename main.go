/*
	CSCI4030U: Big Data Project Part 1
	Main
	Author: Michael Valdron
	Date: Feb 12, 2018
*/
package main

import "fmt"

// The filename of the data source to be used.
const FNAME = "data/retail.dat"

// The threshold used in Apriori and PCY algorithms.
const THOLD float32 = 0.01

// Print frequent items and itemsets (pairs, triples)
func print(freq_itemsets []map[string]int) {
	for _, item_set := range freq_itemsets {
		fmt.Println(item_set)
	}
}

func main() {
	/*
	   Stores the frequent itemsets as hash maps where
	   the key is the itemset and the value is the frequency
	   of the itemset.
	*/
	var freq_itemsets []map[string]int

	// Gets frequent itemsets from the Apriori algorithm
	freq_itemsets = Apriori(FNAME, THOLD)

	// Print Apriori results
	print(freq_itemsets)

	// Gets frequent itemsets from the PCY algorithm
	//freq_itemsets = PCY(FNAME, THOLD)

	// Print PCY results
	//print(freq_itemsets)
}
