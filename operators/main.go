package main

import (
	"fmt"
	"math"
)

func main() {
	// multiplication
	var radius = 12.0
	area := math.Pi * radius * radius
	fmt.Println("area is", area)

	// division
	half := 1 / 2
	fmt.Println("half with integer division", half)

	halfFloat := 1.0 / 2.0
	fmt.Println("half with float division", halfFloat)

	//squaring
	badThreeSquared := 3 ^ 2
	fmt.Println("bad Three squared", badThreeSquared)

	goodThreeSquared := math.Pow(3.0, 2.0)
	fmt.Println("good Three squared", goodThreeSquared)

	// modulus
	remainder := 50 % 3
	fmt.Println("remainder", remainder)

	// unary operators
	x := 3
	x++
	fmt.Println("x is now", x)

	x--
	x--
	fmt.Println("x is now", x)

	// cant do
	// y := x++

}
