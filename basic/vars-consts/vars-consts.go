/*BASIC-DATA.go

There are two types of storage, var's/variables or consts/constants. The value of a variable can change through the course of a program but the data type cant. The value and data type of a constant cannot change after it has been declared.

There are three basic data types,
String - Stores text characters inside of speech marks
Int - Full name integer, stores a whole number
Bool - Full name boolean, stores a true or false value

There are three primary ways to declare variables, each have different purposes,

var NAME DATATYPE = DATA
This is the most clear method of defining variables but it is also the longest way of defining them. An example is: var Num int = 10

var NAME DATATYPE
Similar to the above method, this is a clear way of defining a variable but there is no value given when the variable is created.

NAME := DATA
This is a common way of declaring a variable as it is quick and easy to write and read. Go is smart enough to use the data that you input and find what type of data it is and then set the variable to that data type.
For example: myName := octocat is the same as var myName string = "octocat"

*/

package main

func main() {
	//Declaring Variables

	//most clear method using the data type int used for storing whole numbers
	var num int = 0
	//most clear method with an undefined value
}
