package big_data_project

import "strings"

func get_unique_items(fcontents string) []int {
	baskets := strings.Split(fcontents, "\n")
	items := []int{}
	for _, b := range baskets {
		for index, item := range strings.Split(b, " ") {
			items[index] = int(item)
		}
	}

	return items
}
