/*
	CSCI4030U: Big Data Project Part 1
	Procedures
	Author: Michael Valdron
	Date: Feb 12, 2018
*/
package main

import (
	"bufio"
	"os"
	"reflect"
	"sort"
	"strings"
)

func readLines(f *os.File) []string {
	var lines []string

	fs := bufio.NewScanner(f)
	for fs.Scan() {
		lines = append(lines, strings.TrimSuffix(fs.Text(), " "))
	}

	return lines
}

func checkIn(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func fillCollecInt(v int, l int) []int {
	my_collec := make([]int, l)
	for i := range my_collec {
		my_collec[i] = v
	}

	return my_collec
}

func getKeyStr(a_map map[string]int) []string {
	var keys []string
	for _, k := range reflect.ValueOf(a_map).MapKeys() {
		keys = append(keys, k.String())
	}
	return keys
}

func getUniqueItems(fcontents []string, min_supp int) map[string]int {
	item_counts := make(map[string]int)
	for _, basket := range fcontents {
		for _, item := range strings.Split(basket, " ") {
			keys := getKeyStr(item_counts)
			if len(keys) == 0 || !checkIn(item, keys) {
				item_counts[item] = 1
			} else {
				item_counts[item] += 1
			}
		}
	}

	for key, value := range item_counts {
		if value < min_supp {
			delete(item_counts, key)
		}
	}

	return item_counts
}

func getBaskets(k_items []string, init_items []string) []string {
	var baskets []string
	for _, k_item_set := range k_items {
		for _, item := range init_items {
			if !checkIn(item, strings.Split(k_item_set, ",")) {
				sel_item_set := append(strings.Split(k_item_set, ","), item)
				sort.Strings(sel_item_set)
				if !checkIn(strings.Join(sel_item_set, ","), baskets) {
					baskets = append(baskets, strings.Join(sel_item_set, ","))
				}
			}
		}
	}

	return baskets
}

func getFreqTuples(f *os.File, tuples []string, min_supp int) map[string]int {
	item_set_counts := make(map[string]int)

	for _, tuple := range tuples {
		f.Seek(0, 0)
		fs := bufio.NewScanner(f)
		for fs.Scan() {
			fline := strings.TrimSuffix(fs.Text(), " ")
			similar_count := 0
			for _, item := range strings.Split(tuple, ",") {
				if checkIn(item, strings.Split(fline, " ")) {
					similar_count += 1
				}
			}
			if similar_count == len(strings.Split(tuple, ",")) {
				keys := getKeyStr(item_set_counts)
				if !checkIn(tuple, keys) {
					item_set_counts[tuple] = 1
				} else {
					item_set_counts[tuple] += 1
				}
			}
		}
	}

	for key, value := range item_set_counts {
		if value < min_supp {
			delete(item_set_counts, key)
		}
	}

	return item_set_counts
}
