package main

import (
	"fmt"
)

// basic types (numbers, strings, booleans)
var myInt int

// Discouraged to use
// var myInt16 int16
// var myInt32 int32
// var myInt64 int64

var myUint uint // Only positives

var myFloat float32
var myFloat64 float64

func main() {
	myInt = 10
	myUint = 200
	myFloat = 10.1
	myFloat64 = 100.1

	// The strings are inmutable
	// we create a new string and assign to the myString variable
	myString := "Trevor"
	fmt.Println(myString)
	myString = "John"

	var myBool = true
	fmt.Println(myBool)
}
