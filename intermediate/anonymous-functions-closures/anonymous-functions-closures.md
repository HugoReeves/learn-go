# Anonymous Functions and Closures

[Code](anonymous-functions-closures.go)

Anonymous Functions allow us to make and use functions inline and Anonymous Functions can create closures. This lesson has a lot to do with the scope which you will learn about in a following lesson.

### Anonymous Functions
Anonymous functions allow us to declare an use functions without naming them. This helps to minimize the number of things within the scope. Anonymous functions are often declared as return values of other functions, this keeps them anonymous and allows them to run themselves automatically which is why the function can remain nameless.

### Closure

Closure is often seen as a very complicated topic for those not familiar with the subject. Closure allows you to create functions that retain a certain values across runs. In the code for this lesson, we define a function that returns a function that returns an int. The int increments itself when the function runs.
```go
func countBy5() func() int {
  count := 0
  return func() int {
    count += 1
    return count*5
  }
}
```
Then, we add the following code to func main().
```go
by5 := countBy5()

fmt.Println(by5())
fmt.Println(by5())
fmt.Println(by5())
```
Essentially by5 is a variable attached to the function countBy5, this way, when we call the by5() function, the function countBy5() runs. So, when we print by5(), the first time we print the number 5, but due to closure each subsequent calling of the function is incremented by five. So,
```
5
10
15
```
Is printed.

Due to the nature of the scope, the variables that we initialize in a function defined outside of func main() cannot be accessed inside func main().
```go
func closure() func() string {
	storeThisVal := "Closure"
	return func() string {
		return storeThisVal
	}
}

func main() {
	storeThisVal := "Not closure"
	close := closure()

	fmt.Println(close())	//Closure
	fmt.Println(storeThisVal)	//Not closure
}
```
So, when we run the following code, the variable, storeThisVal, defined inside of the function closure() is not the same or accessable inside of func main().

### Next

* [Home](../../README.md)
* [Intermediate Overview](../intermediate.md)
* [Scope](../scope/scope.md)
