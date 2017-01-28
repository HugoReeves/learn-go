# Methods

[Code](../methods.go)

Interfaces allow us to group types together based on the methods available for each type. This is known as polymorphism and it's an important part of creating complex code that uses multiple different types of structs, while keeping our code simple.

### Setup

Before we begin learning about interfaces, we should define a few structs and methods first. In this example, we will be have three structs, each themed around business.
```go
type boss struct {
	fName            string
	rank             string
	fixedSalary      int
	contractDuration int
}

type deskWorker struct {
	fName            string
	promotionPending bool
	fixedSalary      int
	contractDuration int
}

type investor struct {
	fName       string
	lName       string
	netWorth    int
	trustworthy bool
}
```
We then define the same method for both the boss and deskWorker structs. It is important to note that name of the method and the return value is the same. In this example, the only difference between the two methods is the recieving type, deskWorker and boss.
```go
func (p deskWorker) contractCost() int {
	return p.fixedSalary * p.contractDuration
}

func (p boss) contractCost() int {
	return p.fixedSalary * p.contractDuration
}
```
Now that we have three types and two simmilar methods setup, we can begin learning about interfaces and polymorphism.

### How an Interface works, and how to define an Interface.

It is important to understand how interfaces work, before we can make our own. Interfaces work implicitly rather than explicitly, meaning that instead of listing which types to include in an interface, we just list the methods a type needs to have to be a part of the interface. The most simple way to read an interface is "If a type has the following methods, that return the following values, it is a member of this interface".
Interfaces are laid out simmilar to structs, using the type statement. Standard syntax for creating an interface is the following.
```go
type INTERFACENAME interface {
	METHOD() RETURNTYPE
	...
	METHOD() RETURNTYPE
}
```
* The interface name is how we refer to the interface.
* Each method listed is a method that a type requires to be considered part of the interface.
* The RETURNTYPE states the return value of the method listed.

We have already defined two methods with the same name and return type, now, we can add them to an interface.
```go
type worker interface {
	contractCost() int
}
```
Using our statement from earlier, we can say, "If a type has the contractCost method, and returns the type int, it is a member of the worker interface". So, the boss and deskWorker types both have the contractCost method, therefore they are members of the worker interface.
Now that we have our interface defined, we can create a function that takes any member of the worker interface as an arg.
```go
func printCost(p worker) {
	fmt.Print("The total contract cost is ", p.contractCost(), "\n")
}
```
This way, we can pass a boss or a deskWorker to the printCost function, and they will both act the same. This is polymorphism, and it will become more useful to you as you progress to more complex projects.

### Next

* [Home](../../README.md)
* [Intermediate Overview](../intermediate.md)
* [Next Lesson: Panic, Recover and Defer](../panic-recover-defer/panic-recover-defer.md)