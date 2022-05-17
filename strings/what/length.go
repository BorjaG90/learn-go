package main

import "fmt"

func length() {
	//             012345678901234567890123456789012345
	courseName := "Learn Go for Begginners Crash Course"

	for i := 0; i <= 21; i++ {
		fmt.Println(string(courseName[i]))
	}
	fmt.Println()

	fmt.Println("Lengt of courseName is", len(courseName))

	var mySlice []string
	mySlice = append(mySlice, "one")
	mySlice = append(mySlice, "two")
	mySlice = append(mySlice, "three")

	fmt.Println("my slice has", len(mySlice), "elements")
	fmt.Println("The last element in mySlice is", mySlice[len(mySlice)-1])
}
