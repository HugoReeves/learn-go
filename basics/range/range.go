package main

import "fmt"

func main() {
	//Slices and Arrays
	words := []string{"Learn", "golang", "now!"}
	var message string
	for _, word := range words {
		message += word + " "
	}
	fmt.Println(message)

	//Maps
	kvs := map[string]string{"name": "DevelopDevs", "location": "global"}
	for key, value := range kvs {
		fmt.Println(key, "is the key for the value", value)
	}

	//Strings
	for index, unicode := range "Hi" {
		fmt.Println(index, unicode)
	}
}
