package main

import "fmt"

func boolprint() {
	// it is about bool and print types
	var val bool = false

	fmt.Printf("My data type %T \n", val)

	// different types of print statements
	//  print , println,printf, sprintf
	fmt.Print("it just print a line  --------->")
	fmt.Println("it print a line and cursor goes to next line")
	fmt.Println("i am the new cursor")

	// this println is like python println
	fmt.Println(" My value is ", val, "and i am boolean")

	fmt.Printf("I am new line with formating ... i can be used with variables %v", val) // it doesn't move the cursor to next line
	//
	fmt.Print("\n")
	// differet format specifiers
	number := 34
	str := "i am cool"
	fmt.Printf("My age is %v  and %v \n", number, str) // it puts the value
	fmt.Printf("My age is %q  and %q \n", number, str) // it puts the double quotes
	fmt.Printf("My age is %T  and %T \n", number, str) // it prints the data types

	// there are more format specifiers but these are most used ones

	// sprintf others are also there sprintln , sprint but not much used but same concept
	//  it saves the formated string

	var str1 = fmt.Sprintf("My age is %T  and %T \n", number, str)
	fmt.Println(str1)

}
