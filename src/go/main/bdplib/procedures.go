package bdplib

// Debug
import (
	"strconv"
	"strings"
)

func get_unique_items(fcontents string) []int {
	baskets := strings.Split(fcontents, "\n")
	items := []int{}
	for i, b := range baskets {
		for j, item := range strings.Split(b, " ") {
			items[i+j], _ = strconv.Atoi(item)
		}
	}

	return items
}
