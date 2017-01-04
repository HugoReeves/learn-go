#Range

[Code](range.go)

Range allows us to iterate over values in multiple different datatypes.

###The basics

Range usually outputs two values, we can catch 1 or both of these values. To catch both values the syntax is `V1CAUGHT, V2CAUGHT := range VALUE`. You can also ignore one of the values by using the standard underscore `_` character. The following code utilizes a slice with words and creates a message by iterating through the values in the slice and then prints the message. Because the index of each value is not needed, the underscore is used to not store the index value. This code could easily be adapted to take a slice of names and make a message listing the names in a realistic fashion.
```go
words := []string{"Learn", "golang", "now!"}
var message string
for _, word := range words {
    message += word + " "
}
fmt.Println(message)
```

###Datatype specific interactions

The following explains the specific interactions range has with different datatypes.

* Slices and Arrays - As seen in the previous example, range used on a slice or array returns a first value of the index and a second value at the index location.

* Maps - The first value of the range is the key and the second value is the value associated with the key.

* Strings - Range performed on strings returns the index of a character and then the Unicode for the character.

###Next

* [Home](../../README.md)
* [Basics Overview](../basics.md)
* [Maps](../maps/maps.md)
