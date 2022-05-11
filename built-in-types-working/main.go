package main

import "fmt"

// basic types (numbers, strings, booleans)
/* var myInt int

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
	log.Println(myString)
	myString = "John"

	var myBool = true
	log.Println(myBool)
} */

// aggregate types (array, struct)

type Car struct {
	NumberOfTires int
	Luxury        bool
	BucketSeats   bool
	Make          string
	Model         string
	Year          int
}

func main() {
	// array
	/* var myStrings [3]string
	myStrings[0] = "cat"
	myStrings[1] = "dog"
	myStrings[2] = "fish"

	fmt.Println("First element in array is", myStrings[0]) */

	//struct
	/* var myCar Car
	myCar.NumberOfTires = 4
	myCar.Luxury = false
	myCar.Make = "Seat" */
	myCar := Car{
		NumberOfTires: 4,
		Luxury:        true,
		BucketSeats:   true,
		Make:          "Bugatti",
		Model:         "Veron",
		Year:          2021,
	}

	fmt.Printf("My car is a %d %s %s", myCar.Year, myCar.Make, myCar.Model)
}

// reference types (pointers, slices, maps, functions, channels)

// interface type
