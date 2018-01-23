package big_data_project

import (
	"fmt"
	"io/ioutil"
)

const fname = "../data/test.dat"

func main() {
	f, err := ioutil.ReadFile(fname)

	if err != nil {
		fmt.Printf("Error reading file %s.\n", fname)
	}

	items := get_unique_items(string(f))

	fmt.Printf("%d", items[0][0])

}
