package main

import (
	"fmt"
	"strings"
)

func main() {

	newString := "Go is a great programming language. Go for it!"

	if strings.Contains(newString, "Go") {
		//newString = strings.Replace(newString, "Go", "Golang", 1)
		newString = strings.ReplaceAll(newString, "Go", "Golang")
	}

	fmt.Println(newString)

	// string comparison
	a := "A"
	if a == "A" {
		fmt.Println("a is equal to A")
	}

	if "A" > "B" {
		fmt.Println("A is greater than B")
	} else {
		fmt.Println("A is not greater than B")
	}

	if "Alpha" > "Absolute" {
		fmt.Println("Alpha is greater than Absolute")
	} else {
		fmt.Println("Alpha is not greater than Absolute")
	}

	badEmail := " me@here.com  "
	badEmail = strings.TrimSpace(badEmail)

	fmt.Printf("=%s=", badEmail)
	fmt.Println()
}
