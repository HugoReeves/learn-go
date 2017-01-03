package main

import "fmt"

func main() {
	//Creating a map
	myMap := make(map[string]bool)
	fmt.Println("Created a map:", myMap) //map[]

	//Adding entries
	myMap["workDone"] = true
	myMap["projectComplete"] = true
	myMap["gitPushed"] = false
	fmt.Println("Entered values:", myMap) //map[workDone:true projectComplete:true gitPushed:false]

	//Getting values from keys
	fmt.Println("Value from key:", myMap["gitPushed"])

	//Deleting keys
	delete(myMap, "gitPushed")
	fmt.Println("Deleted key and value gitPushed:", myMap) //map[workDone:true projectComplete:true]

	//Declaring and initializing
	newMap := map[string]int{"age": 21, "children": 0, "siblings": 3}
	fmt.Println("Declared and initialized newMap in one statement:", newMap) //map[age:21 children:0 siblings:3]
}
