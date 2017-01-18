# Structs

[Code](structs.go)

Structs are organised lists of data with defined datatypes, you can easily organise 'objects' in your code with structs. Structs are useful because you can define their format once, and then create any number of varaibles or entries based on the single datatype. Structs are commonly used in OOP (Object Oriented Programming), which is a current method of programming used to organise your code.

### Defining a Struct

Defining structs follows a pretty basic syntax, you enter the name for your struct and then the fields and datatypes associated with each field. Structs you define are considered types, or datatypes, this means you can use them wherever a type such as int or string is accepted. You can make funcs that accept a struct or use a struct as a type definition for an array. The syntax to define a struct is as follows,
```go
type STRUCTNAME struct {
    fieldname datatype
    fieldname datatype
    ...
    fieldname datatype
}
```
For example:
```go
type enemy struct {
    id int
    health int
    race string
}
```

### Using Structs

Now that we have defined a struct, how do we interact wih structs?

1. Recording Data as a variable - You can record data implicitly by supplying all the entries to the different fields sepparated by commas. Using our enemy struct for example, entering this code, `ralph := enemy{"Ralph", 1, 20, "Goblin"}`, creates a variable with the enemy type and values name: "Ralph", id: 1, health: 20, type: "Goblin". As seen in the code, we can then call this variable and print it with the following output, `{Ralph 1 20 Goblin}`.
    You can also record data in a more explicit manor, by stating the field for each value. For example, `james := enemy{name: "James", health: 10, race: "Humanoid"}`, creates a variable with the enemy type that states `{James 0 10 Humanoid}` when printed. It is important to note here, that the id value was not stated, therefore, the id will be zero valued, as is evident in the print out.

1. Recording data into a slice/array/etc. - Since structs are treated as datatypes, we can make a slice that stores the enemy struct type. The following code initializes a slice with 3 empty values that stores entries of the enemy data type, `enemies := make([]enemy, 3)`. Now we can add to the slice our two enemy variables that we created earlier,
    ```go
    enemies[0] = ralph
	enemies[1] = james
    ```
    And we can also add a new value directly to the slice,
    ```go 
    enemies[2] = enemy{"Tim", 2, 40, "Alien"}
    ```
    Now, we have 3 enemy type entries in the enemies slice. We can print them with the following code `fmt.Println(enemies)` and recieve the output, `[{Ralph 1 20 Goblin} {James 0 10 Humanoid} {Tim 2 40 Alien}]`

1. Accessing values - You can easily access values corresponding to a field name with the syntax `STRUCT.FIELDNAME`. Using our code, we can access the health of the ralph enemy struct using the following code `ralph.health`. So `fmt.Println(ralph.health)` prints, 20, which is the health value of the enemy struct that is the ralph variable.

### Next

* [Home](../../README.md)
* [Intermediate Overview](../intermediate.md)
* [Next Lesson: Methods](../methods/methods.md)
