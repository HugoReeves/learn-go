package main

import "fmt"

func main() {
	//Creating empty arrays
	//Empty int array
	var myIntArray [5]int
	fmt.Println(myIntArray) //[0 0 0 0 0]

	//Empty string array
	var myStringArray [5]string
	fmt.Println(myStringArray) //[     ]

	//Empty boolean array
	var myBoolArray [5]bool
	fmt.Println(myBoolArray) //[false false false false false]

	//Creating an array with my own data
	commonNames := [3]string{"John", "Jeff", "Joe"}
	fmt.Println(commonNames) //[John, Jeff, Joe]

	//Creating multiple dimension arrays
	var twoDimension [2][3]int
	fmt.Println(twoDimension)

	var threeDimension [2][3][4]int
	fmt.Println(threeDimension)

	//Indexing arrays
	//Printing and updating an individual index
	fmt.Println(commonNames[2]) //Joe
	commonNames[2] = "Ben"
	fmt.Println(commonNames[2]) //Ben

	//Indexing between indexes
	fmt.Println(commonNames[0:2]) //John Jeff

	//indexing above indexes
	fmt.Println(commonNames[0:]) //John Jeff Ben

	//Indexing below indexes
	fmt.Println(commonNames[:2]) //John Jeff
}
