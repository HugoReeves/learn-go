package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Global Scope - accessable everywhere
	globalScope := "accessable everywhere"

	//Scope within a statement, for, func, etc.
	for i := 0; i < 2; i++ {
		fmt.Println("Global Scope:", globalScope)
		fmt.Println("Values defined inside curly braces, accessable inside curly braces", strconv.Itoa(i))
	}
	//Trying to call i would not work now because i does not exist within this scope

	//Seperate scopes can have variables of the same name
	for scopeExample := 1; scopeExample < 10; scopeExample++ {
		//This is weird syntax because we end up incrementing scopeExample by two each time loop runs
		//scopeExample can have one value here
		scopeExample++
	}
	//If we tried to call scope example here it wouldn't work because it is not in the scope
	//But when we make a variable scopeExample in this scope, it can be called
	scopeExample := "Can be called"
	fmt.Println("Scope example:", scopeExample)
}
