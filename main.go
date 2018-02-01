/*
	CSCI4030U: Big Data Project Part 1
	Main
	Author: Michael Valdron
	Date: Feb 12, 2018
*/
package main

import "fmt"

const FNAME = "data/test.dat"
const THOLD = 10

func main() {
	var freq_item_sets []map[string]int
	freq_item_sets = Apriori(FNAME, THOLD)
	fmt.Println(freq_item_sets[0])
	fmt.Println(freq_item_sets[1])
	fmt.Println(freq_item_sets[2])
}
