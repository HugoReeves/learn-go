#Operators and Comparators

[Code](./operators-comparators.go)
We can conduct basic math and logic operations using different operators. These allow us to perform actions with values. We can also use comparators to compare things.

###Math Operators
There are five primary math operators, each have different purposes.

* Addition `+`
	The addition operator allows us to add together numbers.
	For example: `5+2` is equal to `7`

* Subtraction `-`
	The subtraction operator allows us to subtract numbers.
	For example: `5-2` is equal to `3`

* Multiplication `*`
	The multiplication operator allows us to multiply numbers.
	For example: `5*2` is equal to `10`

* Division `/`
	The divisional operator allows us to divide numbers. However unlike the previous operators, division is slightly more complex. The division will return the greatest number or times it can be divided with a remainder.
	For example: `10/2` is equal to `5` but `9/2` is equal to `4` and 9 can be divided completely only four times by the number two.

* Modulus `%`
	Following on from the divisional operator, the modulus operator allows us to find the remainder of a division.
	For example: `10%2` is equal to `0` as there is no remainder but `9%2` is equal to `1` as the remainder of this division is 1.

###Logic Operators
Logic operators allow us to make decisions based on data. There are three common logic operators. The use of these will make more sense in the next section, but for now we will use booleans to show how Logic Operators work.

* OR `||`
	The OR operator allows us check if one thing or another is true.
	For example: `true || false` is equal to `true` because one of the values is true.

* AND `&&`
	The AND operator allows us to check if two or more things are true.
	For example: `true && false` is false as not all of the values are true.

* NOT `!`
	The NOT operator allows us to check what something is not.
	For example: `!true` is equal to `false` and `!false` is equal to `true`

###Comparators
Comparators allow us to compare values and if they meet the criteria, return false. This will make more sense in the next section using If Else statements.

* is equal to `==`
	Is true if x is equal to y.
	For example: `10 == 5` is equal to `false`

* is not equal to `!=`
	Is true if x is not equal to y.
	For example: `10 != 5` is equal to `true`

* is greater than or equal to `>=`
	Is true if x is greater than or equal to y.
	For example: `10 >= 5` is equal to `true`

* is greater than `<`
	Is true if x is greater than y.
	For example: `10 > 5` is equal to `true`

* is less than or equal to `>=`
	Is true if x is less than or equal to y.
	For example: `10 <= 5` is equal to `false`

* is less than `<`
	Is true if x is less than y.
	For example: `10 < 5` is equal to `false`


###Next

* [Home](../../README.md)
* [Basics Overview](../basics.md)
* [Next: If Else](../if-else/if-else.md)
