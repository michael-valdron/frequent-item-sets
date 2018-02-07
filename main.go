/*
	CSCI4030U: Big Data Project Part 1
	Main
	Author: Michael Valdron
	Date: Feb 12, 2018
*/
package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

const APRIORI_FLAG = "a"
const PCY_FLAG = "p"
const PERCENT_MAX = 100.0
const PERCENT_MIN = 0.0

type Itemset struct {
	tuple []int
	count int
}

type Itemsets []Itemset

func (sets Itemsets) Len() int               { return len(sets) }
func (sets Itemsets) Less(i int, j int) bool { return sets[i].count < sets[j].count }
func (sets Itemsets) Swap(i int, j int)      { sets[i], sets[j] = sets[j], sets[i] }

func mergeSets(freq_itemsets []map[int][]int, freq_itemset_counts []map[int]int) []Itemsets {
	itemsets_and_counts := []Itemsets{}
	for k, itemsets := range freq_itemsets {
		itemsets_and_counts = append(itemsets_and_counts, Itemsets{})
		for key, value := range itemsets {
			itemsets_and_counts[k] = append(itemsets_and_counts[k], Itemset{value, freq_itemset_counts[k][key]})
		}
		sort.Sort(sort.Reverse(itemsets_and_counts[k]))
	}
	return itemsets_and_counts
}

// Print frequent items and itemsets (pairs, triples)
func print(title string, itemsets_and_counts []Itemsets) {
	fmt.Printf("\n%s\n", title)
	for i := 0; i < len(title); i++ {
		fmt.Print("=")
	}
	fmt.Println("\n(itemsets: counts)")
	for k, tuples := range itemsets_and_counts {
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
		for _, itemset := range tuples {
			print_str := ""
			for i, item := range itemset.tuple {
				if (i + 1) < len(itemset.tuple) {
					print_str += strconv.Itoa(item) + ", "
				} else {
					print_str += strconv.Itoa(item)
				}
			}
			fmt.Printf("%s: %d\n", print_str, itemset.count)
		}
		switch k {
		case 0:
			fmt.Printf("Number of Items: %d\n", len(tuples))
		case 1:
			fmt.Printf("Number of Pairs: %d\n", len(tuples))
		case 2:
			fmt.Printf("Number of Triples: %d\n", len(tuples))
		default:
			fmt.Printf("Number of Tuples: %d\n", len(tuples))
		}
	}
	fmt.Println()
}

func main() {
	/*
	   Stores the frequent itemsets as hash maps where
	   the key is the itemset and the value is the frequency
	   of the itemset.
	*/
	var fname *string  // The filename of the data source to be used.
	var thold *float64 // The threshold used in Apriori and PCY algorithms.
	var freq_itemsets []map[int][]int
	var freq_itemset_counts []map[int]int
	var start_time time.Time
	var finish_time time.Time

	fname = flag.String("f", "", "Filename of the selected dataset.")
	thold = flag.Float64("t", 0, "Percent threshold of frequent itemset counts.")
	algorithm := flag.String("alg", APRIORI_FLAG, "Specify which algorithm to run.  a - Apriori, p - PCY")

	flag.Parse()

	if *fname != "" && (*thold > PERCENT_MIN && *thold < PERCENT_MAX) {
		if *algorithm == APRIORI_FLAG {
			fmt.Println("Apriori started.")
			fmt.Println("Please wait..")
			start_time = time.Now()
			// Gets frequent itemsets from the Apriori algorithm
			freq_itemsets, freq_itemset_counts = Apriori(*fname, float32(*thold/PERCENT_MAX))
			finish_time = time.Now()
			fmt.Println("Printing result..")
			// Print results
			print("Apriori Results", mergeSets(freq_itemsets, freq_itemset_counts))
			fmt.Println("Done!")
			fmt.Println("Program Finished.")
			fmt.Printf("Execution took %d minutes and %d seconds.\n",
				int(finish_time.Sub(start_time).Minutes()),
				(int(finish_time.Sub(start_time).Seconds()) % 60))
		} else if *algorithm == PCY_FLAG {
			fmt.Println("PCY started.")
			fmt.Println("Please wait..")
			start_time = time.Now()
			// Gets frequent itemsets from the PCY algorithm
			//freq_itemsets, freq_itemset_counts = PCY(*fname, *thold)
			finish_time = time.Now()
			fmt.Println("Printing result..")
			// Print results
			print("PCY Results", mergeSets(freq_itemsets, freq_itemset_counts))
			fmt.Println("Done!")
			fmt.Println("Program Finished.")
			fmt.Printf("Execution took %d minutes and %d seconds.\n",
				int(finish_time.Sub(start_time).Minutes()),
				(int(finish_time.Sub(start_time).Seconds()) % 60))
		} else {
			fmt.Printf("Please enter either %s or %s under the alg flag.\n", APRIORI_FLAG, PCY_FLAG)
			os.Exit(1)
		}

	} else {
		if *fname == "" {
			fmt.Println("Please enter a filename to select a dataset to start the program.")
		}
		if *thold > PERCENT_MIN && *thold < PERCENT_MAX {
			fmt.Printf("Please enter a valid Percent Threshold less than %f%% and greater than %f%%.",
				PERCENT_MIN, PERCENT_MAX)
		}
		os.Exit(1)
	}
}
