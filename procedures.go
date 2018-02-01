/*
	CSCI4030U: Big Data Project Part 1
	Procedures
	Author: Michael Valdron
	Date: Feb 12, 2018
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

func openFile(fname string) (*os.File, bool) {
	f, err := os.Open(fname)
	if err != nil {
		fmt.Printf("Error reading file %s.\n", fname)
		return f, false
	} else {
		return f, true
	}
}

func getFileLength(fname string) int {
	f, is_open := openFile(fname)
	defer f.Close()
	if is_open {
		fs := bufio.NewScanner(f)
		line_c := 0

		for fs.Scan() {
			line_c += 1
		}

		return line_c
	} else {
		return 0
	}
}

/*
   Checks if a value is in a array/slice. The values passed
   must be of the type string.
*/
func checkIn(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

/*
   Creates an entire array filled with a specific
   value and length passed.  This is an integer
   version of the algorithm.
*/
func fillCollecInt(v int, l int) []int {
	my_collec := make([]int, l)
	for i := range my_collec {
		my_collec[i] = v
	}

	return my_collec
}

/*
   Gets a slice of all the keys which a passed
   hash map contains.  This is the string version
   of the algorithm.
*/
func getKeyStr(a_map map[string]int) []string {
	var keys []string
	for _, k := range reflect.ValueOf(a_map).MapKeys() {
		keys = append(keys, k.String())
	}
	return keys
}

func filterItemsets(itemset_counts map[string]int, min_supp int) {
	for key, value := range itemset_counts {
		if value < min_supp {
			delete(itemset_counts, key)
		}
	}
}
