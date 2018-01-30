/*
	CSCI4030U: Big Data Project Part 1
	Main
	Author: Michael Valdron
	Date: Jan 29, 2018
*/
package main

import "github.com/CSCI4030U-Project/src/go/main/bdplib"

const FNAME = "../../../data/test.dat"
const THOLD = 5

func main() {
	bdplib.Apriori(FNAME, THOLD)
}
