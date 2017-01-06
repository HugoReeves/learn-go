package main

import "fmt"

func main() {
	fmt.Println("Is", 17, "prime?", checkPrime(17))   //true
	fmt.Println("Is", 15, "prime?", checkPrime(15))   //false
	fmt.Println("Is", 239, "prime?", checkPrime(239)) //true
}

//Here we define the checkprime function
//checkPrime is a function that takes one int argument called input and returns a single boolean
func checkPrime(input int) bool {
	//we create the factorCount variable
	factorCount := 0
	//We loop through all the numbers from 1 to the input
	for i := 1; i <= input; i++ {
		//If the input is perfectly divisible by i, i is a factor of the input. This could easily be changed to add the value of i, which is a factor of the input to a slice which stores all the factors of the input.
		if input%i == 0 {
			factorCount++
		}
	}
	//if there are only two factors of the input, the input is prime so the checkprime function should return the bool true. This could also be done using the len() of a slice if you decided to add the factors of the input to a slice.
	if factorCount == 2 {
		return true
	} else {
		//Otherwise return false
		return false
	}
}
