package main

import "fmt"

func main() {
	fmt.Println("Three argument for loop, for i := 1; i <= 30; 1++")
	for i := 1; i <= 30; i++ {
		fmt.Println(i)
	}

	fmt.Println("One argument for loop, for i < 10")
	i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println("No argument for loop")
	j := 1
	for {
		fmt.Println(j)
		if j == 10 {
			break
		}
		j++
	}
}
