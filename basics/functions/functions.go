package main

import "fmt"

func bingo(userI int, bingoNum int) bool {
	if userI == bingoNum {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(bingo(7, 238))	//false
}
