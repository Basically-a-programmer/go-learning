package main

import "fmt"

// Variable declaration outside the main
var val int = 10

func main() {
	// Printing the outside variable
	fmt.Println("outside variable:", val)

	// Variable with the same name but inside the main function
	var val float64 = 10.74
	fmt.Println("inside the main function:", val)

	// Inside the if statement
	if true {
		// Variable with the same name but inside the if block
		var val string = "i am string"
		fmt.Println("inside the if:", val)
	}
}
