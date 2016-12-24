package main

import "fmt"

func main() {
	//Declaring Variables
	//The basic method
	var num int = 10

	//Declaring multiple values at once
	var textOne, textTwo string = "Hello", "There"

	//Defined with no value
	var decision bool

	//The shorthand and common method
	myName := "octocat"

	//Here we can see what these variables are equal to
	fmt.Println("The following are the variables printed.")
	fmt.Println("num =", num)                                               //10
	fmt.Println("textOne and textTwo printed together =", textOne, textTwo) //Hello There
	fmt.Println("decision =", decision)                                     // If not defined, defaults to false
	fmt.Println("myName =", myName)                                         //octocat

	//Declaring constants
	//The basic method
	const constantString string = "Constant String"

	//A slightly shorter method where the type is inferred
	const constantInteger = 10

	//Here we can see what these constant are equal to
	fmt.Println("The following are the constants printed.")
	fmt.Println("constantString =", constantString)
	fmt.Println("constantInteger =", constantInteger)
}
