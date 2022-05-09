package main

import (
	"fmt"
	"starting/packageone"
)

var one = "One"

func main() {
	var something = "This is a block level variable"
	fmt.Println(something)
	myFunc()

	newString := packageone.PublicVar
	fmt.Println("From packageone:", newString)

	// secondString := packageone.privateVar
	// fmt.Println("From packageone:", secondString)

	packageone.Exported()

}

func myFunc() {
	var one = "the number one"
	fmt.Println(one)
}
