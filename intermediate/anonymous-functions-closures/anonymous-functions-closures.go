package main

import "fmt"

func countBy5() func() int {
  count := 0
  return func() int {
    count += 1
    return count*5
  }
}

func closure() func() string {
	storeThisVal := "Closure"
	return func() string {
		return storeThisVal
	}
}

func main() {
	//Count up by 5
	by5 := countBy5()

	fmt.Println(by5())
	fmt.Println(by5())
	fmt.Println(by5())

	//The vars are seperate due to the scope
	storeThisVal := "Not closure"
	close := closure()

	fmt.Println(close())	//Closure
	fmt.Println(storeThisVal)	//Not closure
}
