package main

import (
	"fmt"
)

//The factorial function takes an int and returns nothing
func factorial(origin *int) {
	//Input is designated as the value of the pointer origin
	value := *origin

	//The counter for the loop, i, is set equal to input
	i := value

	//The loop runs until the counter, i, is equal to 1
	for {
		if i == 1 {
			break
		}

		//Make the input equal to itself * i
		value = value * i

		//Decrement i by 1
		i--
	}

	//Update the value at the pointer of origin to be equal to the value of input
	*origin = value
}

func main() {
	//Initialize and check the myNumber variable
	myNumber := 10
	fmt.Println("myNumber before being passed through the factorial() function:", myNumber) //10

	//The pointer of the myNumber variable is passed to the factorial function, this will change it's value to be equal to the factorial
	factorial(&myNumber)

	//Print the new value of myNumber
	fmt.Println("myNumber after being passed through the factorial() function:", myNumber) //36288000
}
