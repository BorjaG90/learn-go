package main

import (
	"starting/packageone"
)

// variables
// declare a package level variable for the main
// package named myVar
var myVar = "This is package level variable,"

func main() {
	// declare a block variable for the main function
	// called blockVar
	var blockVar = "this is block level variable,"

	// declare a package levele variable in the packageone
	// package named PackageVar

	// create an exported function in packageone called PrintMe

	// in the main function, print out the values of myVar
	// blockVar, and PackageVar on one line using the PrintMe
	// function in packageone
	packageone.PrintMe(myVar, blockVar)
}
