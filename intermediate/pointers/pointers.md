# Pointers

[Code](pointers.go)

Pointers allow us to pass the memory location of variables. This allows us to do lots of cool stuff, like changing variables from different scopes in functions (You will learn about that soon). This is important as it allows us to temporarily access variables in scopes that we would not usually be able to access.

When we create a variable, `myString := "The original message"`, it is given a location in memory. We can call this location with a pointer. In order to get the pointer to the variable myString, an and/& symbol should be appended.
For example:
```go
myString := "The original message"
fmt.Println("Pointer to myString:", &myString)
```
Will print the pointer to the myString variable. It is important to note that even when we change the value of a variable, it retains the same location, therefore all pointers will still work no matter the value of the variable.

Now that we know how to reference pointers with the `&NAME` syntax, we can create functions that take pointers as arguments. We can create a basic function, `func changeString(args) returntype { code }`. In the arguments section we can add an assigned name, in this case input and then state that its expected datatype is `*string`, where the star character indicates the type is a pointer and string indicates that it is a pointer to a string. Now we have the code `func changeString(input *string) returntype { code }`, we can get rid of the return type as we will not be returning anything with this function. In order to change the value associated with the pointer, we refernece *input and then assign it a value of the string datatype, `*input = "The string has been changed"`. Now, we have a function that takes a pointer to a string type variable, and changes it's value to "The string has been changed". Our final code looks like this,
```go
func changeString(input *string) {
	*input = "The string has been changed"
}
```

Go see the [code](pointers.go) to see how this works in action.

### Next

* [Home](../../README.md)
* [Intermediate Overview](../intermediate.md)
* [Challenge 3: Factorial Function](../../challenges/intermediate/factorial-function/factorial-function.md)
