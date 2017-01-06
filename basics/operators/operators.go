package main

import "fmt"

func main() {
	//Arithmetic Operations
	//Addition
	fmt.Println("5 + 2 =", 5+2) //7

	//Subtraction
	fmt.Println("5 - 2 =", 5-2) //3

	//Multiplication
	fmt.Println("5 * 2 =", 5*2) //10

	//Division
	fmt.Println("10 / 2 =", 10/2) //5
	fmt.Println("9 / 2 =", 9/2)   //4

	//Modulus
	fmt.Println("10 percent 2 =", 10%2) //0
	fmt.Println("9 percent 2 =", 9%2)   //1

	//Logic Operations
	//OR
	fmt.Println("true || false =", true || false) //true

	//AND
	fmt.Println("true && false =", true && false) //true

	//NOT
	fmt.Println("!true =", !true)   //false
	fmt.Println("!false =", !false) //true

	//Comparison Operations will be used in the next section, if-else.
	//is equal to
	fmt.Println("10 == 2 =", 10 == 2) //false

	//is not equal to
	fmt.Println("10 != 2 =", 10 != 2) //true

	//is greater than or equal to
	fmt.Println("10 >= 2 =", 10 >= 2) //true

	//is greater than
	fmt.Println("10 > 2 =", 10 > 2) //true

	//is less than or equal to
	fmt.Println("10 <= 2 =", 10 <= 2) //false

	//is less than
	fmt.Println("10 < 2 =", 10 < 2) //false

}
