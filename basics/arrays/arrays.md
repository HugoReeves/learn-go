#Arrays

[Code](arrays.go)
Arrays have a set length and allow us to store individual pieces of data, of the same type in a single variable. When you initialize an array in Go, you can initialize the array with predefined values, or create the array with empty values. Importantly, the length of an array cannot be changed, in the next lesson we will learn about slices, which are similar to arrays but the length of a slice can change.

###Creating Arrays
There are multiple ways that we can create arrays. You can think about arrays as lists but arrays are lists with a defined length. It is also important to note that when we create arrays, we must define what data type the array can store, int, string, etc.

* Creating an empty array of specified length - Being able to initialize an array with empty values allows us to build an array that we can fill in with our own data as our program runs. Creating an empty array is similar to creating an empty variable. We create arrays by writing in our program, `var NAME [ARRAY LENGTH]DATATYPE`

	For example: `var myArray [10]int` creates an array with the name myArray that can store 10 ints.

* Creating an array with data - Creating an array with your own data can be useful if you want to make a constant array or you already have all the data you need to start the array. To create a variable array with your own data use the syntax `NAME := [ARRAY LENGTH]DATATYPE{value1, value2, ..., valueN}`

	For example: `commonNames := [3]string{"John", "Jeff", "Joe"}`

* Creating dimensional arrays - Dimensional arrays are arrays within arrays. A two dimensional array is an array that contains an array inside of it. We can create two dimensional arrays by using the following syntax `var NAME [DIMENSION 1 LENGTH][DIMENSION 2 LENGTH]int`.

	For example: `var threeDimension [2][3][4]int`

###Indexing Arrays
Now that we can create empty arrays and arrays with our own data, lets look at how we interact with arrays. Indexing is how we change or look at certain values of our array. The syntax for indexing an array is `ARRAYNAME[INDEX]`, it is important to note that the first entry in our array is at the 0 index, and therefore if we create an array of length 5, the final entry in the array is at index 4. This is a little confusing but it is standard across many programming languages.

* Indexing Single Values - Once we know how to index an array, we can update or print individual values just like any other variable.

	For example: `fmt.Println(commonNames[2])` will print `Joe` which is the final name in the array. Similarly commonNames[2] = "Ben" updates the entry at the second index to be equal to Ben.

* Indexing between indexes - We are able to index between two indexes by using the following syntax `ARRAYNAME[INDEX1:INDEX2]`. With this syntax, we are selecting everything above and including index 1 but excluding index 2.

	For example: `fmt.Println(commonNames[0:2])` will print values 1 and two which in this case are John and Jeff.

* Indexing below and above indexes - Similarly, you can index below and above indexes by using the : character.

	For example: `fmt.Println(commonNames[0:])` will print all values including and above the 0 index, in this case John, Jeff and Ben.
	`fmt.Println(commonNames[:2])` will print all values below but excluding the 2 index, in this case John and Jeff.

###Next

* [Home](../../README.md)
* [Basics Overview](../basics.md)
* [Unknown](../../challenges/basics/20-even-odd/20-even-odd.md)
