/*
	CSCI4030U: Big Data Project Part 1
	Procedures
	Author: Michael Valdron
	Date: Jan 29, 2018
*/
package bdplib

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
)

func readLine(f *os.File) string {
	fs := bufio.NewScanner(f)

	if fs.Scan() {
		return fs.Text()
	} else {
		return ""
	}
}

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

// Old Code
/*func getBaskets(k_items []string, init_items []string, min_supp int) map[string]int {
	basket_counts := make(map[string]int)
	for _, k_item_set := range k_items {
		for _, item := range init_items {
			sel_item_set := append(strings.Split(k_item_set, ","), item)
			sort.Strings(sel_item_set)
			if len(k_item_set) == 1 || !checkIn(item, strings.Split(k_item_set, ",")) {
				basket_counts[strings.Join(sel_item_set, ",")] = 1
			} else {
				basket_counts[strings.Join(sel_item_set, ",")] += 1
				fmt.Println(strings.Join(sel_item_set, ","))
			}
		}
	}

	for key, value := range basket_counts {
		if value < min_supp {
			delete(basket_counts, key)
		}
	}

	return basket_counts
}*/

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

func Apriori(fname string, t_hold int) {
	var k = 0
	var fcontents []string
	var item_set_counts []map[string]int
	var min_supp int

	f, err := os.Open(fname)
	defer f.Close()
	if err != nil {
		fmt.Printf("Error reading file %s.\n", fname)
		return
	}
	fcontents = readLines(f)
	min_supp = t_hold // * len(fcontents)
	item_set_counts = append(item_set_counts, getUniqueItems(fcontents, min_supp))
	fcontents = nil

	for len(item_set_counts[k]) > 0 && k < 2 {
		k_item_sets := getKeyStr(item_set_counts[k])
		init_item_sets := getKeyStr(item_set_counts[0])
		candidate_item_set := getBaskets(k_item_sets, init_item_sets)
		item_set_counts = append(item_set_counts, getFreqTuples(f, candidate_item_set, min_supp))
		k += 1
	}

	fmt.Println(item_set_counts[0])
	fmt.Println(item_set_counts[1])
	fmt.Println(item_set_counts[2])
}
