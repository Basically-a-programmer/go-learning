package main

import "fmt"

func hello() {
	// types of variable
	//  bool , string
	// int , uint (unsigned int) , float , complex , byte, rune ( char)
	//  you can have auto interpreter for variable declaration using

	total := 0 // auto declare

	average, message := 23.6, "it's the average" // multiple declaration
	var num int = int(average)
	const pul int = 40 // constant
	total = num + pul
	fmt.Println("total", total)
	fmt.Println("the average", average)
	fmt.Println("the display", message)
	fmt.Println("sum = ", num+pul)

}
