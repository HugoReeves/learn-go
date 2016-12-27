package main

import "fmt"

func main() {
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			//if i is even, because an even number/2 has a remainder of 0
			fmt.Println(i, "- EVEN")
		} else {
			//otherwise i must be odd
			fmt.Println(i, "- ODD")
		}
	}
}
