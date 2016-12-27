#If Else Statements

[Code](./if-else.go)
If else statements allow programs to make decisions based on values that are given to the statements. The basis of if else statements is a set of values that interact using operators and comparators, the interaction of these operators and comparators must be true for the statement to execute. You can chain together statements using if, if else and else, but only one statement can run.

###Types of statements
There are three primary ways that we use to make if else statements.

* if
	Of x is true, do this.
	For example: `if x > 5 {
		return "x is greater than five"
	}`

* else if
	Otherwise, check if this is true and if so, do this.
	For example: `else if x > 2 {
		return "x is greater than 2"
	}`

* else
	Otherwise do this.
	For example: `else {
		return x
	}`

If we write out these statements in order, we get the following.
```go
if x > 5 {
	fmt.Println("x is greater than five")
} else if x > 2 {
	fmt.Println("x is greater than 2")
} else {
	fmt.Println(x)
}
```
Reading this code through, we get. If x is greater than five, print `x is greater than five`, and if that is not true. Check if x is greater than 2, if so, print `x is greater than 2`, and if that is not true, print x. With this set of statements, when know that if x=7 `x is greater than five will be printed` and if x is equal to or less than 2, x will be printed.

###Next

* [Home](../../README.md)
* [Basics Overview](../basics.md)
* [Next: For Loops](../for/for.md)
