package main

import "fmt"

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
