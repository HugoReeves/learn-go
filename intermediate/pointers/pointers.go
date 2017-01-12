package main

import (
	"fmt"
)

//Func that changes a string
func changeString(input *string) {
	*input = "The string has been changed"
}

func main() {
	//Define the string here
	myString := "The original message"
	//Printing the location of the variable
	fmt.Println("Pointer to myString:", &myString)
	//Printing the value of the variable
	fmt.Println("Before passing myString through changeString function, myString:", myString)

	//Pass it through the function and then check it's value
	changeString(&myString)
	fmt.Println("After passing myString through changeString function, myString:", myString)
	//The pointer does not change
	fmt.Println("Pointer to myString:", &myString, "still the same")
}
