# Check if Prime

In this challenge, you will attempt to make and utilize a basic Go function that takes an integer input and returns true or false whether the input is a prime. You can check whatever numbers you want but I will be using the following three and they should return the following values when given to the checkPrime function, 17 - True, 15 - False, 239 - True. I suggest you use a function called chackPrime() or something similar to achieve this, this will help you retain what you have learnt recently in the past few lessons. Very little basic help will be given for this challenge, but if you get stuck or need help you can look at the hints section below or checkout my answer to this question [here](check-prime.go).

### Hints
Look away quick! If you want to complete this challenge all on your own, these hints may spoil the challenge for you.

* You may want to use a for loop to loop between 1 and the number to check.
* A prime number has exactly two factors.
* You may want to store and check the number of factors for the input number.

### A solution, explained

The way I chose to solve this problem was with a function defined outside of func main called checkPrime(). The function is defined in the following manor, `func checkPrime(input int) bool { FUNC }`, like we learnt in the last lesson, this function definition means that the function takes one int called input and returns a single boolean. Inside the curly braces we write out function. First of all, we need to initialize the factorCount variable that will hold the number of factors of the input, `factorCount := 0`.
The method I used to find out if a number is prime is by looping through the values of an index, where the index starts at 1 and ends at the input value, 1 <= i <= input. To make this for loop we add inside of our function the code `for i := 1; i <= input; i++ { LOOP }`. Now that we have a loop with index i looping from 1 to the input value, we can add our logic. To find if a given number is a factor of another number, we can use modulus and check if the remainder of a division is 0. If we find a factor, we should increment the factorCount variable by 1. We can use an if else statement such as the following to achieve this logic, `if input%i == 0 { factorCount++ }`.
Finally, outside of our for loop we should add the logic that returns true or false based on the number of factors. As I said earlier, if an input has only two factors, it is a prime otherwise it is not a prime, this should be pretty easy to convert to if else logic.
```go
if factorCount == 2 {
	return true
} else {
	return false
}
```
Now that we have finished our function, we can add this simple code to func main and we are done.
```go
fmt.Println("Is", 17, "prime?", checkPrime(17))
fmt.Println("Is", 15, "prime?", checkPrime(15))
fmt.Println("Is", 239, "prime?", checkPrime(239))
```
Check out the completed and annotated [code](check-prime.go).


### Next

* [Home](../../README.md)
* [Basics Overview](../basics.md)
* [Next lesson: Intermediate Overview](../../../intermediate/intermediate.md)
