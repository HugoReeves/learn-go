# Switches

[Code](switches.go)

Switches are similar to if else statements, and they are actually capable of the exact same tasks as if else statements. The basis of a switch is taking a value, usually stored as a variable and performing logical actions based on the value of the variable.

### Using Switches
There are multiple ways to use switches, but generally you include a variable in the switch statement and the values are referenced in the case statements.

* Basic switch and cases - We use the following syntax for a basic switch and case,
	```go
	switch VARIABLE {
	case VALUE:
		ACTION
	...
	}
	```

	For example:
	```go
	a := 2
	fmt.Print("The number is ")
	switch a {
	case 1:
		fmt.Println("Uno")
	case 2:
		fmt.Println("Dos")
	case 3:
		fmt.Println("Tres")
	}
	```
	Will result in `The number is Dos` being printed.

* Switch as a replacement for an if else statement - We can use switches as a replacement for if else statements. It is important to note that the first case to run will end the switch and no further cases will run. Furthermore, the default case functions like the else case. If no case runs, the default will run.

	For example:
	```go
	b := 12
	switch {
	case b < 10:
		fmt.Println("Less than 10")
	case b > 10:
		fmt.Println("Greater than 10")
	default:
		fmt.Println("The variable is not an int")
	}
	```
	Will print `Greater than 10`, if however b was not an int, the default case would print `The variable is not an int`.

### Next

* [Home](../../README.md)
* [Basics Overview](../basics.md)
* [Maps](../maps/maps.md)
