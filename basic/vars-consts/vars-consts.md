#Variables and Constants
Here we learn about the two ways in which we store data.
Data can be stored as a variable or a constant. Variables are declared once, and their value can change through the course of a program however the data type can't. The value and data type of a constant cannot change after it has been declared.

###Basic Data Types
* String - Stores text characters inside of speech marks
* Int - Full name integer, stores a whole number
* Bool - Full name boolean, stores a true or false value

###Declaring Variables
There are three primary ways to declare variables, each have different purposes,

* `var NAME DATATYPE = DATA`
 This is the most clear method of defining variables but it is also the longest way of defining them.
 An example is: `var num int = 10`

* `var NAME DATATYPE`
 Similar to the above method, this is a clear way of defining a variable but there is no value given when the variable is created.
 For example: `var decision bool`

* NAME := DATA
 This is a common way of declaring a variable as it is quick and easy to write and read. Go is smart enough to use the data that you input and find  what type of data it is and then set the variable to that data type.
 For example: `myName := octocat` is the same as `var myName string = "octocat"`
