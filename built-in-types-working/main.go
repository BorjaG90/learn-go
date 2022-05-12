package main

import (
	"fmt"
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
/*func main() {
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
}*/

// maps
/*
func main() {
	intMap := make(map[string]int)
	// un mapa siempre se pasa por referencia

	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4
	intMap["five"] = 5

	for key, value := range intMap {
		fmt.Println(key, value)
	}

	delete(intMap, "four")

	for key, value := range intMap {
		fmt.Println(key, value)
	}

	el, ok := intMap["four"]
	if ok {
		fmt.Println(el, "is in map")
	} else {
		fmt.Println(el, "isn`t in map")
	}

	intMap["two"] = 4

}
*/

type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

// Funcion con a receiver
func (a *Animal) Says() {
	fmt.Printf("A %s says %s", a.Name, a.Sound)
	fmt.Println()
}

func (a *Animal) HowManyLegs() {
	fmt.Printf("A %s has %d legs", a.Name, a.NumberOfLegs)
	fmt.Println()
}

func main() {
	/* z := addTwoNumbers(2, 4)
	fmt.Println(z) */

	/* myTotal := sumMany(2, 3, 4, 5, 92, 7, -5)
	fmt.Println(myTotal) */

	var dog Animal
	dog.Name = "dog"
	dog.Sound = "guau"
	dog.NumberOfLegs = 4
	dog.Says()

	cat := Animal{
		Name:         "cat",
		Sound:        "miau",
		NumberOfLegs: 4,
	}

	cat.Says()
	cat.HowManyLegs()
}

func addTwoNumbers(x, y int) (sum int) {
	sum = x + y
	return
}

// variatic function
func sumMany(nums ...int) int {
	total := 0
	for _, x := range nums {
		total = total + x
	}

	return total
}
