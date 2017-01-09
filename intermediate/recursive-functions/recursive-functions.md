#Recursive Functions

Recursive Functions are functions that call themselves. Recursive functions work by returning themselves unless a certain if else statement runs.

###Defining Recursive Functions
The basis of a recursive function is a function that returns itself. This is best explained through an example.
```go
package main

import "fmt"

func endlessPrint(num) {
	fmt.Println(num)

	return endlessPrint(num)
}

func main() {
	endlessPrint(3)
}
```
The previous example is not a function you would ever want to use, as it will never end. The function simply repeatedly prints the input value. A much better example is one in which we adjust the value passed to the recursive function each time.
```go
package main

import "fmt"

func endlessPrint(num int) int {
	if num == 0 {
		return 0
	}
	fmt.Println(num)
	num--
	return endlessPrint(num)
}

func main() {
	endlessPrint(10)
}

```
This simply prints the values 10 -> 1 to the console. Although this could be easily achieved using a for loop, it proves the concept of recursive functions.

###Next

* [Home](../../README.md)
* [Intermediate Overview](../intermediate.md)
* [Unknown](../recursive-functions/recursive-functions.md)
