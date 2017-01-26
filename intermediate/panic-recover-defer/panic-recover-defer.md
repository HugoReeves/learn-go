# Panic, Recover and Defer

[Code](panic-recover-defer.go)

Panic, Recover and Defer allow us to handle errors within our program, and understand where errors are occuring.

### Basic Overview

1. Panic - Panic immediatley stops a function from running. We can stop func main from running with a panic statement, but usually, it is far more useful to use panic inside other functions. Panic takes a single string arg, that can be recalled with the recover function. The syntax to write a panic statement is `panic(String)`. Panic is usually used in switches or if else statements.
	
    For example:
    ```go
	for i := 1; i <= 10; i++ {
		if i == 6 {
			panic("Woah, something went wrong.")
		}
		fmt.Println(i)
	}
	```

1. Recover - Recover returns the string from the panic function. Since you pass a string to the panic function, you use recover to return the string from the panic function. In order to use recover, you usually need to pair it with defer. The syntax is simply `recover()`

1. Defer - Defer allows an action to be run at the end of a function, no matter where the defer statement is. The syntax for defer is `defer ACTION`
	
    For example:
    ```go
	func showcaseDefer() {
		defer fmt.Println("We're done counting")
		for i := 1; i <= 10; i++ {
			fmt.Println(i)
		}
	}
	```

### Putting it all together

We can create a function that showcases these three things together. We will want to start this function with a defered recover statement. This is important as if we dont defer the recover, it will not run because a panic statement will stop any code running. Once we have a defered recover statement, we can write whatever code we want, and then add a panic statement. For example:

```go
func showcasePanicRecover() {
	defer func() {
		fmt.Println(recover())
	}()
	for i := 1; i <= 10; i++ {
		if i == 6 {
			panic("Woah, something went wrong.")
		}
		fmt.Println(i)
	}
}
```

### Next

* [Home](../../README.md)
* [Intermediate Overview](../intermediate.md)
* [Next Lesson: Panic, Recover and Defer](../panic-recover-defer/panic-recover-defer.md)