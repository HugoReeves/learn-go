package main

import "fmt"

func main() {
	showcaseDefer()
	showcasePanicRecover()
}

//This function showcases the way defer works
func showcaseDefer() {
	//With this defer clause, we're done counting will be printed once the for loop stops running
	defer fmt.Println("We're done counting")

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

//This function showcases how we use defer, panic and recover together
func showcasePanicRecover() {
	defer func() {
		//By defering this func, we know that even if a panic takes place, we will still print the recover statement
		fmt.Println(recover())
	}()
	for i := 1; i <= 10; i++ {
		//Because panic runs before the println, only the numbers up to 5 will be printed.
		if i == 6 {
			//In this case the panic message will be recovered by the defered recover statement
			panic("Woah, something went wrong.")
		}
		fmt.Println(i)
	}
}
