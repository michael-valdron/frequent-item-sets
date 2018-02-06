/*
	CSCI4030U: Big Data Project Part 1
	Main
	Author: Michael Valdron
	Date: Feb 12, 2018
*/
package main

import (
	"fmt"
	"time"
)

type Itemset struct {
	tuple []int
	count int
}

// The filename of the data source to be used.
const FNAME string = "retail.dat"

// The threshold used in Apriori and PCY algorithms.
const THOLD float32 = 0.01

// Print frequent items and itemsets (pairs, triples)
func print(title string, freq_itemsets []map[int][]int, freq_itemset_counts []map[int]int) {
	fmt.Println(title)
	for i := 0; i < len(title); i++ {
		fmt.Print("=")
	}
	fmt.Println("\n(itemsets: counts)")
	for k, itemsets_counts := range freq_itemsets {
		switch k {
		case 0:
			fmt.Println("\nFrequent Items:")
		case 1:
			fmt.Println("\nPairs:")
		case 2:
			fmt.Println("\nTriples:")
		default:
			fmt.Println("\nOther Itemsets:")
		}
		for items, counts := range itemsets_counts {
			fmt.Printf("%s: %d\n", items, counts)
		}
	}
}

func main() {
	/*
	   Stores the frequent itemsets as hash maps where
	   the key is the itemset and the value is the frequency
	   of the itemset.
	*/
	var freq_itemsets []map[int][]int
	var freq_itemset_counts []map[int]int
	var start_time time.Time
	var finish_time time.Time

	fmt.Println("Apriori started.")
	fmt.Println("Please wait..")
	start_time = time.Now()
	// Gets frequent itemsets from the Apriori algorithm
	// freq_itemsets, freq_itemset_counts = Apriori(FNAME, THOLD)
	Apriori(FNAME, THOLD)
	finish_time = time.Now()
	fmt.Printf("Done!  Took %d minutes and %d seconds.\n",
		int(finish_time.Sub(start_time).Minutes()),
		(int(finish_time.Sub(start_time).Seconds()) % 60))
	fmt.Println("Printing result..")
	// Print Apriori results
	print("Apriori Results", freq_itemsets, freq_itemset_counts)
	fmt.Println("Done!")

	// fmt.Println("PCY started.")
	// fmt.Println("Please wait..")
	// start_time = time.Now()
	// // Gets frequent itemsets from the PCY algorithm
	// freq_itemsets = PCY(FNAME, THOLD)
	// finish_time = time.Now()
	// fmt.Printf("Done!  Took %d minutes and %d seconds.\n",
	// 	int(finish_time.Sub(start_time).Minutes()),
	// 	(int(finish_time.Sub(start_time).Seconds()) % 60))
	// fmt.Println("Printing result..")
	// // Print PCY results
	// print("PCY Results", freq_itemsets)
	// fmt.Println("Done!")
	fmt.Println("Program Finished.")
}
