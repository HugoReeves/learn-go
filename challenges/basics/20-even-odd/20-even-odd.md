#1 to 20 Even and Odd

###Introduction

After learning about the different value types, variables and constants, if else statements and for loops, it's about time for your first challenge! In this challenge you need to use all the skills you have gained, to print out a list of numbers, 1-20 or any other list, with the value even or odd based on the number. When you run your program, it should output something similar to this.
```
1 - ODD
2 - EVEN
3 - ODD
...
19 - ODD
20 - EVEN
```
Now, what are you waiting for? Create a new file either here and call it myanswer.go and get going. The solution to this problem can be found in [20-even-odd.go](20-even-odd.go).
###Hints

1. You may need to look at the [Operators and Comparators](../../../basics/operators-comparators/operators-comparators.md) section to figure out how to find even and odd numbers.

###A Solution, Explained
[Solution](./20-even-odd.go). This is only one way of solving the problem, and as with most programming challenges, there are multiple ways to go about solving the same problem. In this method, and all other methods, we start with the basic outline of a Go program.
```go
package main

import fmt

func main() {
	YOUR CODE GOES HERE
}
```
Now, thinking about this problem, it seems the best way to go about finding a solution is with a for loop that will loop between 1 and 20. So to create this loop we can use a three argument for loop. We will start be intialising an index variable i with the value of 1 `i := 1`. Next, we can make sure that the loop runs so long as i is equal to or less than 20, `i <= 20`, remember, between each argument a ; character is needed. Finally, the third argument is to increment i by one each time the loop runs, to do this, we add as our third argument `i++` which is shorthand for `i = i + 1`.

Now that we have our for loop, we have to implement logic into the loop so that if the number is odd, a certain piece of code runs and if the number is even a different piece of code is run. The primary way of implementing logic and decision making in Go and most other languages is by using if else statements. So in our code we make a statement `if THESE STATEMENTS ARE TRUE { RUN THIS CODE }`. In order to check if a number is even, we can use i%2 which will find the remainder of the division i/2. In the if statement we can write `if i%2 == 0 { I MUST BE EVEN }` this works because if you divide a number by 2, and the remainder is 0, the number is even. Now we can fill in the code block, and make it print the value of i and state whether it is even or odd. `if i%2 == 2 { fmt.Println(i, "- EVEN") }`. This will tell us when a number is even but we still need to print when a number is odd, to do this we can use an else statement because if i is not even it must be odd. Our final code looks a little something like this.
```go
package main

import "fmt"

func main() {
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Println(i, "- EVEN")
		} else {
			fmt.Println(i, "- ODD")
		}
	}
}
```
You can find the code [here](./20-even-odd.go).

###Next

* [Home](../../README.md)
* [Basics Overview](../basics.md)
* [Next: Unknown](../../basics/unknown/logic-math.md)
