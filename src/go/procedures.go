package big_data_project

// Debug
import "fmt"

import (
	"strings"
	//"strconv"
)

func get_unique_items(fcontents string) [][]int {
	baskets := strings.Split(fcontents, "\n")
	items := [][]int{{}}
	for _, b := range baskets {
		for _, item := range strings.Split(b, " ") {
			fmt.Printf("%s", item)
			//items[i][j] = strconv.Atoi(item)
		}
	}

	return items
}
