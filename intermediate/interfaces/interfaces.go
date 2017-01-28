package main

import "fmt"

//Define the worker interface
type worker interface {
	contractCost() int
}

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

//The following two methods are identical except for the type they recieve
func (p deskWorker) contractCost() int {
	return p.fixedSalary * p.contractDuration
}

func (p boss) contractCost() int {
	return p.fixedSalary * p.contractDuration
}

//Here, we create a function that recieves anything of the worker interface type
func printCost(p worker) {
	fmt.Print("The total contract cost is ", p.contractCost(), "\n")
}

func main() {
	//defining a deskWorker and a boss, these both fall under the worker interface
	johnny := deskWorker{"Johnny", false, 45000, 2}
	ted := boss{"Ted", "Overseer", 75000, 1}

	//We define bill as an investor
	bill := investor{"Bill", "West", 999999, false}

	//Printing all the people
	fmt.Println(johnny)
	fmt.Println(ted)
	fmt.Println(bill)

	//Use the printCost func that takes an arg worker
	//This is interesting as johnny and ted are of different types but they both have the same interface of worker
	//Therefore they can both be used as args to the printCost function
	printCost(johnny)
	printCost(ted)

	//We cannot use printCost on bill because he is an investor and does not fall under the worker interface
	//printCost(bill) would result in an error
}
