package bdplib

// Debug
import (
	"strconv"
	"strings"
)

func checkIn(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func GetUniqueItems(fcontents []string) []int {
	var items []int
	for _, b := range fcontents {
		for _, item := range strings.Split(b, " ") {
			a, _ := strconv.Atoi(item)
			if len(items) == 0 || !checkIn(a, items) {
				items = append(items, a)
			}
		}
	}

	return items
}
