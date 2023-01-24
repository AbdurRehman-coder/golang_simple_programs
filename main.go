package main

import (
	"fmt"
)

func main() {
	// my first line of code in go, just printing
	fmt.Println("My first go programm")

	/// Declaration with initialization
	/* const a int8 = 34
	var b int8 = 9
	var result = a + b
	// formatted print
	fmt.Printf("the result is: %d\n", result)
	// simple print
	fmt.Println("the result is: ", result)
	*/

	// var name = "Abdur Rehman"
	// userName := "Shinwari"

	// fmt.Println("My name is:", name, userName)

	/// Declaration without initialization
	var i int
	var j string
	i = 20
	j = "without initialization"
	fmt.Println("i: ", i)
	fmt.Println("j: ", j)

	/// Declaration of multiple variables
	var fname, lname = "Arshad", "Ali"
	fmt.Println("first & last name: ", fname, lname)

	aa, b, c := 5, 10, 55
	fmt.Println("a = ", aa, " b = ", b, " c = ", c)

	/// Short Variable Declaration in Golang
	name := "Johne Doe"
	age := 24
	fmt.Println(name, " ", age)

	/// Zero Values
	var quantity float32 // it will print 0 because the zero value of [float32] is zero
	fmt.Println(quantity)
	var race int8
	fmt.Println(race)
	var isEligible bool // zero value of bool is false
	fmt.Println(isEligible)

	var number uint8 = 255 // it will give error on 256 because it is greate then range, it also give error on negative number
	fmt.Println(number)

	typeChecking := "working hard"
	fmt.Printf("type is: %T\n", typeChecking)
}
