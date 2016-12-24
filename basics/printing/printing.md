#Printing

[Code](./printing.go)
As this is the first lesson there will be some basics to go over. Printing is one of the most basic operations in any functional language. Go has multiple ways to print and these work through the fmt package, fmt is the standard go package that offers higher level functionality ontop of Go's largely low level operation.

###The outline of a Go program
Every go program is centered around a package main with other packages that are imported to perform functions and add functionality. package main is where all the functions of a program are called and it acts as a hub for your program. When we are starting off, we will write all of our code in this package main, but as you move to more advanced programs you will start to separate your code into different files. So we start with package main, then we import our other packages and then we have a function main which is the primary function that runs. Here is the outline of our code in it's most basic form.
```go
package main

import "fmt"

func main() {

}
```
When we are starting off, we will write all of our code inside this package main.

###Basic Types of Printing
There are two basic types of printing, Print and Println. These are both functions of the fmt package. We can print multiple things in the same statement by separating them with commas. We use backslash n inside of a printed string to print a new line, this is not needed at the end of a Println as a Println automatically creates a new line.
1. Print
	`fmt.Print()` will print whatever is stated. We can print multiple values in the same statement by using commas. You need to manually add spaces between different values with this method.
	For example: `fmt.Print("Hello World \n")` will print `Hello World` followed by a new line and `fmt.Print("Number ", 7)` will print `Number 7. Notice how the space after the word number must be explicitly included. If we do not include a space, `fmt.Print("Number", 3)` the statement `Number3` will be printed.
1. Println
	`fmt.Println()` will print whatever is stated but each value will be automatically separated by a space. This will also start a new line at the end of the statement.
	For example: `fmt.Println("Hi", "there")` will print `Hi there`. Notice how a space was included between the two words, this is because the comma used to input multiple values adds a space automatically.

###Next

* [Home](../../README.md)
* [Basics Overview](../basics.md)
* [Next: Variables and Constants](../values/values.md)
