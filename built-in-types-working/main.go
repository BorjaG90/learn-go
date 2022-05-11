package main

import "fmt"

// reference types (pointers, slices, maps, functions, channels)

func main() {
	x := 10

	myFirstPointer := &x

	fmt.Println("x is", x)
	fmt.Println("myFirstPointer is", myFirstPointer)

	*myFirstPointer = 15
	fmt.Println("x is now", x)

	changeValueOfPointer(&x)
	fmt.Println("After function called x is now", x)

	// interface type
}

func changeValueOfPointer(num *int) {
	*num = 25
}
