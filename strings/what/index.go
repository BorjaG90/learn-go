package main

import "fmt"

func index() {
	//             012345678901234567890123456789012345
	courseName := "Learn Go for Begginners Crash Course"

	fmt.Println(courseName[0])
	fmt.Println(courseName[6])

	for i := 0; i <= 21; i++ {
		fmt.Println(string(courseName[i]))
	}
	fmt.Println()

	for i := 13; i <= 21; i++ {
		fmt.Println(string(courseName[i]))
	}
	fmt.Println()

}
