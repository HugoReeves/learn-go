package main

import "fmt"

func main() {
	//Creating slices
	sliceA := make([]int, 3)
	fmt.Println(sliceA) //[0 0 0]

	//appending
	sliceA = append(sliceA, 5)
	fmt.Println(sliceA) //[0 0 0 5]

	sliceA = append(sliceA, 6, 7)
	fmt.Println(sliceA) //[0 0 0 5 6 7]

	//len()
	fmt.Println(len(sliceA)) //6

}
