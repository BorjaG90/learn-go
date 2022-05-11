package main

import (
	"fmt"
	"sort"
)

// reference types (pointers, slices, maps, functions, channels)

// pointers
/*
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
*/

// slices
func main() {
	var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "fish")
	animals = append(animals, "cat")
	animals = append(animals, "horse")

	fmt.Println(animals)

	//for _, animal := range animals {
	for i, animal := range animals {
		fmt.Println(i, animal)
	}

	fmt.Println("Element 0 is", animals[0])
	fmt.Println("First two elements are", animals[0:2])
	fmt.Println("The slice is", len(animals), "elements long")

	// sort
	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))

	sort.Strings(animals)
	fmt.Println("Sorted?", sort.StringsAreSorted(animals))
	fmt.Println(animals)

	animals = DeleteFromSlice(animals, 1)
	fmt.Println(animals)

}

func DeleteFromSlice(a []string, i int) []string {
	a[i] = a[len(a)-1]
	a[len(a)-1] = ""
	a = a[:len(a)-1]
	return a
}
