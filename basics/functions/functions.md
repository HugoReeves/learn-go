#Functions

[Code](functions.go)

Functions are a building block for the programs we make, and they can be found in all functional programming languages. Functions allow us to create separated pieces of code that allow us to easily perform repetitive functions throughout our code. Functions are relatively basic things but there are some more advanced ways in which we can use functions, we will learn more about functions in the next lesson.

###Defining functions

The functions we will be creating in this lesson take a certain number of arguments with defined types, and return a single value with a defined type. Defining a basic function follows a simple syntax, `func NAME(argname argdatatype, ..., argname argdatatype) returndatatype { CODE }`. The argument names are important as they allow us to call the values given to the function and use them as if they are variables.
An example function:
```go
func bingo(userI int, bingoNum int) bool {
	if userI == bingoNum {
		return true
	} else {
		return false
	}
}
```
This function will return true if the users number is equal to the bingo number, or false if the users number is not equal to the bingo number.

###Calling functions

Calling functions uses a syntax similar to most other languages. The syntax is `FUNCNAME(ARG1, ARG2, ...)`. When you call a basic function, it returns a value, in most cases, like when you use a function inside of print statement, the return value of the function will be printed.
For example: `fmt.Println(bingo(7,238))` will print false.

###Next

* [Home](../../README.md)
* [Basics Overview](../basics.md)
* [Intermediate](../../intermediate/intermediate.md)
