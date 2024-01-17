package main

import "fmt"

// Variable declaration outside the main
var val int = 10

// Global variable and can be used globally
// using 2 steps  1 -> import main
//  main.variable_name
// avoid using global variable
var Gloablvar int = 15

func scope() {
	// Printing the outside variable
	fmt.Println("outside variable:", val)

	// inital value
	var i int
	fmt.Println("initial value of the variable", i)

	// Variable with the same name but inside the main function
	var val float64 = 10.74
	fmt.Println("inside the main function:", val)

	// Inside the if statement
	if true {
		// Variable with the same name but inside the if block
		var val string = "i am string"
		fmt.Println("inside the if ->", val)
	}
}
