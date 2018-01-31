/*
	CSCI4030U: Big Data Project Part 1
	Apriori
	Author: Michael Valdron
	Date: Feb 12, 2018
*/
package bdplib

import (
	"fmt"
	"os"
)

func Apriori(fname string, t_hold int) []map[string]int {
	var k = 0
	var fcontents []string
	var item_set_counts []map[string]int
	var min_supp int

	f, err := os.Open(fname)
	defer f.Close()
	if err != nil {
		fmt.Printf("Error reading file %s.\n", fname)
		return []map[string]int{}
	}
	fcontents = readLines(f)
	min_supp = t_hold
	item_set_counts = append(item_set_counts, getUniqueItems(fcontents, min_supp))
	fcontents = nil

	for len(item_set_counts[k]) > 0 && k < 2 {
		k_item_sets := getKeyStr(item_set_counts[k])
		init_item_sets := getKeyStr(item_set_counts[0])
		candidate_item_set := getBaskets(k_item_sets, init_item_sets)
		item_set_counts = append(item_set_counts, getFreqTuples(f, candidate_item_set, min_supp))
		k += 1
	}

	return item_set_counts
}
