package big_data_project

import (
	"fmt"
	"io/ioutil"
)

const fname = ""

func main() {
	f, err := ioutil.ReadFile(fname)

	if err != nil {
		fmt.Printf("Error reading file %s.\n", fname)
	}

}
