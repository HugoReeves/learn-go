package main

import "fmt"

func main() {
	a := 2
	b := 12

	//Basic switch
	fmt.Print("The number is ")
	switch a {
	case 1:
		fmt.Println("Uno")
	case 2:
		fmt.Println("Dos")
	case 3:
		fmt.Println("Tres")
	}

	//Switch as an if else replacement
	switch {
	case b < 10:
		fmt.Println("Less than 10")
	case b > 10:
		fmt.Println("Greater than 10")
	default:
		fmt.Println("The variable is not an int")
	}
}
