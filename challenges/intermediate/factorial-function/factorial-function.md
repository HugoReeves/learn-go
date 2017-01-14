#Factorial Function

[A Solution](factorial-function.go)

In this challenge, ou are tasked to create a function, factorial(), that accepts an int pointer and either using recursive functions, for loops or another method, updates the value of the int pointer to the factorial. The outline of the function should be as follows.
```go
func factorial(NAME *int) {
	CODE

	*NAME = THE FACTORIAL
}
```
The print the new value.

In my answer, I will be using the number 10, and attempting to find it's factorial which should be equal to 36288000.

###Hints
Look away quick! If you want to complete this challenge all on your own, these hints may spoil the challenge for you.

* Using a for loop to complete the task is the easiest method.
* You may need to separate the input value from it's pointer.

###A solution, explained

I chose to solve this problem with a function containing a for loop. The basis of the algorithim was the for loop that would perform the following loop a number of times `value = value * i`. I created an outline for the program with a function that accepts a pointer to an int. The code to perform this is `func (origin *int) { CODE }`, the input is called origin because I will be separating the value of origin from the pointer to origin, this makes it easier to write the for loop. In the code section, I added the following, `value := *origin` which assigns the value at the origin pointer to the value variable. Then, I initialised a counter for the no argument for loop, `i := input`, I used the input for the original value of i because I want each loop to decrement i until it is equal to 1. Next, I added the no argument for loop, starting with an if statement that breaks the loop if i == 1. The important statement that makes value = value * i is added, this works because the value can change but i will always be one less each time the loop runs. Finally, i-- decrements i by 1.

The final function looks something like this:
```go
func factorial(origin *int) {
	value := *origin

	i := value

	for {
		if i == 1 {
			break
		}

		value = value * i

		i--
	}
	*origin = value
}
```
And func main() looks like this:
```go
func main() {
	//Initialize and check the myNumber variable
	myNumber := 10
	fmt.Println("myNumber before being passed through the factorial() function:", myNumber) //10

	//The pointer of the myNumber variable is passed to the factorial function, this will change it's value to be equal to the factorial
	factorial(&myNumber)

	//Print the new value of myNumber
	fmt.Println("myNumber after being passed through the factorial() function:", myNumber) //36288000
}
```

###Next

* [Home](../../README.md)
* [Intermediate Overview](../../../intermediate/intermediate.md)
* [Next Lesson: Structs](../../../intermediate/structs/structs.md)
