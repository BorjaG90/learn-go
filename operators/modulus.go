package main

import "fmt"

func modulus() {
	x := 12
	y := 3

	if x%y == 0 {
		fmt.Println(y, "divides exactly into", x)
	} else {
		fmt.Println(y, "does not divides exactly into", x)
	}

	for m := 1; m <= 12; m++ {
		fmt.Println("The motnh after", m, "is", m%12)
	}
}
