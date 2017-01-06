#Values

[Code](./values.go)
There are different kinds of values in Go. The four most basic types of values are strings, ints, floats and bools. It is important that when we perform actions in Go we know what types of values we are working with.

###Basic Value Types
* String - `string`, stores text.
* Integer - `int`, stores a whole number.
* Float - `float32` or `float64`, stores a decimal number.
* Boolean - `bool`, stores true or false.

###Understand what value types you are using
Knowing what value types you are using when writing in Go is important. If we were to perform the following action of adding together two strings with the value of 7, `fmt.Println("7"+"7")` the action would print `77` not `14`.
This occurs because when you add strings together, they are concatenated, in order to print `14` you would need to write `fmt.Println(7+7)` which adds two integers.

###Next

* [Home](../../README.md)
* [Basics Overview](../basics.md)
* [Next: Variables and Constants](../vars-consts/vars-consts.md)
