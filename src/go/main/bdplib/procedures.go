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
	"strings"
)

func readLines(fname string) []string {
	var lines []string
	f, err := os.Open(fname)
	defer f.Close()
	if err == nil {
		fs := bufio.NewScanner(f)
		for fs.Scan() {
			lines = append(lines, fs.Text())
		}
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

/*func getBaskets(fcontents []string) [][]int {
	var baskets = make([][]int, len(fcontents))
	for i, _ := range fcontents {
		baskets[i] = make([]int, 0, len(fcontents))
		for _, item := range strings.Split(fcontents[i], " ") {
			nitem, _ := strconv.Atoi(item)
			baskets[i] = append(baskets[i], nitem)
		}
	}

	return baskets
}*/

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

func Apriori(fname string, t_hold int) {
	//var k = 0
	var fcontents []string
	var item_set_counts []map[string]int
	var min_supp int

	fcontents = readLines(fname)
	min_supp = t_hold // * len(fcontents)
	if len(fcontents) < 1 {
		fmt.Printf("Error reading file %s.\n", fname)
		return
	}
	item_set_counts = append(item_set_counts, getUniqueItems(fcontents, min_supp))
	fcontents = nil

	/* TODO: Create main loop for pairs and triples */
	//for len(item_set_counts[k]) > 0 && k < 2 {

	//}
}
