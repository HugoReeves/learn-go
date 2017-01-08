package main

import "fmt"

func main() {
	//Variadic Function with raw args
	printArgSlice(1, 2, 7)	//[1 2 7]
	//Variadic function with slice as arg
	classSizes := []int{1, 2, 7}
	printArgSlice(classSizes...) //[1 2 7]

	//Function with multiple return values
	add, sub := myMath(1, 6, 3, 4)
	fmt.Println(add)	//14
	fmt.Println(sub)	//-14
}

//Variadic Function
func printArgSlice(args ...int){
	fmt.Println(args)
}


//Function with multiple return values
func myMath(myArgs ...int) (int, int) {
	add := 0
	sub := 0
	for _, arg := range myArgs {
		add += arg
		sub -= arg
	}
	return add, sub
}
