# Slices

[Code](slices.go)

Slices have a changeable length and allow us to store individual pieces of data, of the same type in a single variable. Slices act and index almost exactly like arrays but the length of a slice can change. Slices are more common in Go code compared to arrays because they are more flexible and they have more ways in which you can interact with them.

### Creating Slices
Creating a Slice as a variable is easy. The syntax for creating a slice is `NAME := make([]DATATYPE, LENGTH)`. This creates a slice with a starting length, but the length of the slice can change.
For example: `sliceA := make([]int, 3)` will make a slice with a length of 3 with empty values.


### Interacting with Slices
You can index slices just like you index arrays. There are also other ways to interact with slices such as the `len()` function.

* Appending - Appending to a slice is an important operation, and one you will likely perform often. The syntax for appending is `SLICENAME = append(SLICENAME, VALUETOAPPEND, MOREVALUES, ...)`. The reason we have to update the slice is because the append function simply returns the slice with the appended value, so you need to update the original slice in order to save the appended value.

	For example: `sliceA = append(sliceA, 5)` adds 5 as the last value of sliceA. You can also append multiple values with the following statement, `sliceA = append(sliceA, 6, 7)`

* `len()` function - The `len()` function returns the length of an array or slice. You simply pass the array or slice name into the function and it returns the length

	For example: `fmt.Println(len(sliceA))` prints an int of the length of the slice.

### Next

* [Home](../../README.md)
* [Basics Overview](../basics.md)
* [Switches](../switches/switches.md)
