# Scope

[Code](scope.go)

The scope is how we manage where things can be called. We can define and use a variable in one scope, but it may not e accessible in another. This allows us to use common variable names such as i for index in for loops, without overwriting values of i in other scopes. The scope can sometimes cause issues in your code, but it's actually a good thing because it allows us to worry a bit less about how we name our variables and functions, and just focus on writing code.

### Basics

The basis of the scope is that we are able to control where variables are accessible. Due to the scope, variables are only accessable inside certain areas. For example, the following variable i as declared in this for statement `for i := 0; i < 2; i++ {}` can only be accessed within that for statement. You can also have variables with the same name in different scopes.
For example:
```go
for scopeExample := 1; scopeExample < 10; scopeExample++ {
	scopeExample++
	//scopeExample has one value here
}
//but we have to redefine it here
scopeExample := "Can be called"
fmt.Println("Scope example:", scopeExample)
```

### Next

* [Home](../../README.md)
* [Intermediate Overview](../intermediate.md)
* [Pointers](../pointers/pointers.md)
