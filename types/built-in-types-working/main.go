package main

import (
	"fmt"
)

// booleans
/* func main() {
	apples := 18
	oranges := 9

	//boleans expression
	fmt.Println(apples == oranges)
	fmt.Println(apples != oranges)

	// > <
	fmt.Printf("%d > %d: %t", apples, oranges, apples > oranges)
	fmt.Println()
	fmt.Printf("%d < %d: %t", apples, oranges, apples < oranges)
	fmt.Println()
	fmt.Printf("%d >= %d: %t", apples, oranges, apples >= oranges)
	fmt.Println()
	fmt.Printf("%d <= %d: %t", apples, oranges, apples <= oranges)
	fmt.Println()
} */

// compound booleans
type Employee struct {
	Name     string
	Age      int
	Salary   int
	FullTime bool
}

func main() {
	jack := Employee{
		Name:     "Jack Smith",
		Age:      27,
		Salary:   40000,
		FullTime: false,
	}

	jill := Employee{
		Name:     "Jill Jones",
		Age:      33,
		Salary:   60000,
		FullTime: true,
	}

	var employees []Employee
	employees = append(employees, jack)
	employees = append(employees, jill)

	for _, em := range employees {
		if em.Age > 30 {
			fmt.Println(em.Name, "is 30 or older")
		} else {
			fmt.Println(em.Name, "is under 30")
		}

		if em.Age > 30 && em.Salary > 50000 {
			fmt.Println(em.Name, "makes more than 50000 and is over 30")
		} else {
			fmt.Println("Either", em.Name, "makes less than 50000 or is under 30")
		}

		if em.Age > 30 || em.Salary > 50000 {
			fmt.Println(em.Name, "makes more than 50000 or is over 30")
		} else {
			fmt.Println(em.Name, "makes less than 50000 and is under 30")
		}

		if (em.Age > 30 || em.Salary < 50000) && em.FullTime {
			fmt.Println(em.Name, "matches our unclear criteria")
		}
	}
}
