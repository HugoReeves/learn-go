package main

import "fmt"

func main() {
	x := 7
	if x > 5 {
		fmt.Println("x is greater than five")
	} else if x > 2 {
		fmt.Println("x is greater than 2")
	} else {
		fmt.Println(x)
	}

	y := 3

	if y > 5 {
		fmt.Println("y is greater than five")
	} else if y > 2 {
		fmt.Println("y is greater than 2")
	} else {
		fmt.Println(y)
	}

	z := 1

	if z > 5 {
		fmt.Println("z is greater than five")
	} else if z > 2 {
		fmt.Println("z is greater than 2")
	} else {
		fmt.Println("z =", x)
	}
}
