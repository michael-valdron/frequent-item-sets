/*
	CSCI4030U: Big Data Project Part 1
	Set
	Author: Michael Valdron
	Date: Feb 12, 2018
*/
package main

import (
	"sort"
	"strings"
)

type set struct {
	set map[string]bool
}

func ArrayToSet(list []string) set {
	set := set{make(map[string]bool)}

	for _, s := range list {
		set.Add(s)
	}

	return set
}

func SplitToSet(s string, sep string) set {
	return ArrayToSet(strings.Split(s, sep))
}

func (set *set) InSet(s string) bool {
	return set.set[s]
}

func (set *set) Add(s string) bool {
	in_set := set.InSet(s)
	set.set[s] = true
	return !in_set
}

func (set *set) IsSubset(superset []string) bool {
	similar_count := 0

	for i := range superset {
		if similar_count == len(set.set) {
			break
		} else if set.InSet(superset[i]) {
			similar_count += 1
		}
	}

	return similar_count == len(set.set)
}

func (set *set) IsSuperset(subset []string) bool {
	similar_count := 0

	for i := range subset {
		if similar_count == len(subset) {
			break
		} else if set.InSet(subset[i]) {
			similar_count += 1
		}
	}

	return similar_count == len(subset)
}

func (set *set) GetItems() []string {
	var keys []string
	for k, _ := range set.set {
		keys = append(keys, k)
	}
	return keys
}

func (set *set) GetSortedItems() []string {
	sorted_arr := set.GetItems()
	sort.Strings(sorted_arr)
	return sorted_arr
}
