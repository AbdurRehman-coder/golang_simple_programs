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
	fmt.Printf("type is: %T\n", typeChecking) // string type

	/// if Statements
	if age > 18 && age < 25 {
		fmt.Println("You are an adult")
	} else if age > 30 && age < 50 {
		fmt.Println("You are now a man")
	} else {
		fmt.Println("Oh!, You are old enough")
	}

	/// Loops, In Golang we have only for loop
	/* for i := 0; i < 10; i += 2 {
		fmt.Println(i)
	} */

	/// Arrays / List
	/// Assign and Access Values
	var theArray [3]string
	theArray[0] = "Canada"
	theArray[1] = "USA"
	// theArray[2] = "France"
	fmt.Println(theArray)

	///Initializing an Array with an Array Literal
	array1 := [5]int{34, 3, 2, 1}
	var yy [5]int = [5]int{10, 20, 30}
	fmt.Println(array1)
	fmt.Println(yy)

	arrList := [3]string{"foo", "jar"} // List
	arrList[2] = "bar"
	fmt.Println(arrList)

	arr := []int{3, 7} // Slic, This is called slice, because it is a growable list
	fmt.Println("length: ", len(arr))
	fmt.Println("capacity: ", cap(arr))

	// adding new values to the list
	arr = append(arr, 20)
	fmt.Println("length1: ", len(arr))
	fmt.Println("capacity1: ", cap(arr))
}
