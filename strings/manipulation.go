package main

import (
	"fmt"
	"strings"
)

func manipulation() {

	// newString := "Go is a great programming language. Go for it!"

	// if strings.Contains(newString, "Go") {
	// 	//newString = strings.Replace(newString, "Go", "Golang", 1)
	// 	newString = strings.ReplaceAll(newString, "Go", "Golang")
	// }

	// fmt.Println(newString)

	// // string comparison
	// a := "A"
	// if a == "A" {
	// 	fmt.Println("a is equal to A")
	// }

	// if "A" > "B" {
	// 	fmt.Println("A is greater than B")
	// } else {
	// 	fmt.Println("A is not greater than B")
	// }

	// if "Alpha" > "Absolute" {
	// 	fmt.Println("Alpha is greater than Absolute")
	// } else {
	// 	fmt.Println("Alpha is not greater than Absolute")
	// }

	// badEmail := " me@here.com  "
	// badEmail = strings.TrimSpace(badEmail)

	// fmt.Printf("=%s=", badEmail)
	// fmt.Println()

	str := "alpha alpha alpha alpha alpha"
	str = replaceNth(str, "alpha", "beta", 3)
	fmt.Println(str)
}

func replaceNth(s, old, new string, n int) string {
	//index
	i := 0

	for j := 1; j <= n; j++ {
		x := strings.Index(s[i:], old)
		if x < 0 {
			// we didnt find it
			break
		}

		//have found it
		i = 1 + x
		if j == n {
			return s[:i] + new + s[i+len(old):]
		}

		i += len(old)
	}
	return s
}
