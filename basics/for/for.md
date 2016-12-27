#For Loops

[Code](./for.go)
For loops allow your code to repeat a piece of code based on a variable or set parameters. Just like If Else statements, For loops are a key part of many code operations and for loops can be found in many languages.

###Different ways to write For statements
There are multiple ways that we can write for statements.

* The three argument for loop
	The most common type of for loop is a three argument for loop. The first argument initialises a variable that is usually used for counting in the loop, this is often called the index variable. The second argument makes a statement that must be true for the loop to run. The final argument states a change to the index variable that occurs after each run through of the loop.
	For example: `for i := 1; i <= 30; i++ { CODE GOES HERE }` In this example, the code will run 30 times. The loop begins by setting i := 1 and then says that as long as i is less than or equal to 30, run the loop, finally it says that after each run through of the loop increase i by 1.

* The single argument for loop
	If you want more control or just a different way to write for loops, single argument for loops are useful. A single argument for loop only requires a statement that must be true in order to run. In the code section, you should add a statement that changes the value of the index variable so that the code will eventually stop.
	For example: `for i < 10 {CODE GOES HERE}`

* No argument for loops
	No argument for loops allow you to run a loop multiple times, and quit the loop using the `break` statement. A common usage is to iterate through answers to a question and use an if statement so that if you find the correct answer, break the loop. As you move from three argument for loops to no argument loops, you gain more control over how the loop operates, and you are able to create more complex loops.
	For example:
	```go
	j := 1
	for {
		fmt.Println(j)
		if j == 10 {
			break
		}
	}
	```
	This will print the numbers 1 to 10.

###Next

* [Home](../../README.md)
* [Basics Overview](../basics.md)
* [Next: Challenge 1: Count to 20 with Even and Odd](../../challenges/basics/20-even-odd/20-even-odd.md)
