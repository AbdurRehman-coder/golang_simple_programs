package main

import (
	"fmt"
)

func main() {
	// my first line of code in go, just printing
	fmt.Println("My first go programm")

	/// Declaration with initialization
	const a int8 = 34
	var b int8 = 9
	var result = a + b
	// formatted print
	fmt.Printf("the result is: %d\n", result)
	// simple print
	fmt.Println("the result is: ", result)

	// var name = "Abdur Rehman"
	// userName := "Shinwari"

	// fmt.Println("My name is:", name, userName)

	var fname, lname = "Arshad", "Ali"
	fmt.Println("first & last name: ", fname, lname)

}
