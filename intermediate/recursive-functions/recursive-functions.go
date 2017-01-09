package main

import "fmt"

func printDown(num int) int {
	//Once the count reaches 0, the recursion stops. The function will stop at this point so the number 0 will not be printed.
	if num == 0 {
		return 0
	}
	//Print the current value of num, this is the part that prints 10 -> 1
	fmt.Println(num)
	//Decrease num by 1. This is important as just like a for loop that doesnt change the index value, you must change the index in order to make the recursion stop eventually.
	num--
	//This is what makes the recursion occur, because the function returns itself, it loops.
	return printDown(num)
}

func main() {

	fmt.Println("Print 10 -> 1")
	//This will print 10 -> 1
	printDown(10)

	fmt.Println("Print 5 -> 1")
	//Print down from 5
	printDown(5)
}
