#Advanced Functions

[Code](adv-func.go)

Learning how to use advanced features of functions such as Variadic functions and functions with multiple return values is an important step towards creating complex programs.

###Functions with an Arbitrary Number of Arguments (Variadic Functions)

So far we have learnt to create functions with a defined number of inputs, however, sometimes we want to make functions that can take an arbitrary number of inputs or a slice and perform actions on the values.

* Defining Variadic Functions - Defining Variadic Functions is a similar process to defining regular functions. The pattern for naming your arguments is similar but a ... is included before the argument datatype.

	For example: `func printArgSlice(args ...int){ fmt.Println(args) }` defines a function that takes an arbitrary number of ints or a single slice and prints the slice of the inputs.

* Using Variadic Functions - You can use Variadic Functions just like normal functions with the following syntax `FUNCNAME(ARG1, ARG2, as many args as you want)`. Using the example function from above, `printArgSlice(1, 2, 7)` would print `[1 2 7]`. More commonly though, you will be using slices with Variadic Functions. The syntax for using slices with Variadic Functions is as follows, `FUNCNAME(SLICE...)`.

	For example: Using our printArgSlice function, and a slice `classSizes := []int{1, 2, 7}`, the following code `printArgSlice(classSizes...)` will print `[1 2 7]`.

###Functions with Multiple Return Values

Functions that can return multiple values can be useful for many different applications, specifically, you can create complex functions that achieve multiple parts of a problem, and just access certain values when necessary.

* Creating Functions with Multiple Return Values - The syntax for creating a function with multiple return values is, `func name(args) (return1type, return2type)`.

	For example: This example also uses a Variadic Function so keep an eye out for that.
	```go
	func myMath(myArgs ...int) (int, int) {
		add := 0
		sub := 0
		for _, arg := range myArgs {
			add += arg
			sub -= arg
		}
		return add, sub
	}
	```

* Using Functions with Multiple Return Values - Using functions with multiple return values is similar to capturing values from a Range statement. You capture each value into a variable or ignore a value with the underscore, _ character.

	For example: Capturing the values from the previous myMath function and printing them can be done with the following code.
	```go
	add, sub := myMath(1, 6, 3, 4)
	fmt.Println(add)
	fmt.Println(sub)
	```

###Next

* [Home](../../README.md)
* [Basics Overview](../basics.md)
* [Maps](../maps/maps.md)
