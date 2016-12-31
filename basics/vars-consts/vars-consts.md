#Variables and Constants

[Code](./vars-consts.go)
Here we learn about the two ways in which we store values.
Values can be stored as a variable or a constant. Variables are declared once, and their value can change through the course of a program however the type can't. Constants are declared once and they cannot be changed once declared.

###Declaring Variables
There are three primary ways to declare variables, each have different purposes.

* `var NAME TYPE = VALUE`
	This is the most basic method of defining variables but it is also the longest way of defining them.
	An example is: `var num int = 10`

* `var NAME1, NAME2 TYPE = VALUE1, VALUE2`
	In this example we declare multiple variables of the same type.
	An example is: `var textOne, textTwo string = "Hello", "There"`

* `var NAME TYPE`
	Similar to the first method, this is a clear way of defining a variable but there is no value given when the variable is created. Vars created with this will have a zero value until they are changed, so a bool will be false and an int will be 0.
	For example: `var decision bool`

* `NAME := VALUE`
	This is a common way of declaring a variable as it is quick and easy to write and read. Go is smart enough to use the data that you input and infer what type it is.
	For example: `myName := octocat` is the same as `var myName string = "octocat"`

###Declaring Constants
There are two main methods in which constants are declared

* `const NAME TYPE = VALUE`
	This is the most basic way that constants are declared. The name is given, then the type and then the value which must conform to the type.
	For example: `const constantString string = "Constant String"`

* `const NAME = VALUE`
	This is a slightly shorter method in which the type is inferred. The name is given and then a value is given and Go infers the type based on the value.
	For example: `const constantInteger = 10`

###Next

* [Home](../../README.md)
* [Basics Overview](../basics.md)
* [Next: Operators and Comparators](../operators-comparators/operators-comparators.md)
