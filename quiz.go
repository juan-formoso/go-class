package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz game!")

	/*
		VARIABLES AND DATA TYPES
		string => everything in a double quotation mark
		int => whole numbers
		uint => whole numbers with no negative numbers
		float64 => decimal numbers
		bool => true or false
		name := "John" | short way
		%v is a placeholder for the variable
		%d does the same but for decimal numbers

		PRINT AND SCAN
		Scan let the user input data
		\n is a new line

		CONDITIONS
		!= is not equal to
		== is equal to
		< is less than
		> is greater than
		<= is less than or equal to
		>= is greater than or equal to

		CONDITIONALS
		if statement
		if else statement
	*/

	fmt.Printf("Enter\n your name: ")
	var name string
	fmt.Scan(&name)

	fmt.Printf("Hello %v, welcome to the game!", name)

	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("Yay, you are old enough to play! :D")
	} else {
		fmt.Println("Ops! You are not old enough to play! :(")
		return // this will stop the program
	}

	fmt.Printf("What is better, a dog or a cat? ")
	var answer string
	fmt.Scan(&answer)

	fmt.Println(answer)
}
