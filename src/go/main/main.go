package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/CSCI4030U-Project/src/go/main/bdplib"
)

const fname = "../../../data/test.dat"

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

func main() {
	fcontents := readLines(fname)

	if len(fcontents) < 1 {
		fmt.Printf("Error reading file %s.\n", fname)
		return
	}

	items := bdplib.GetUniqueItems(fcontents)

	fmt.Print(items)

}
