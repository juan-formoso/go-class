package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz game!")

	// VARIABLES AND DATA TYPES
	// string => everything in a double quotation mark
	// int => whole numbers
	// uint => whole numbers with no negative numbers
	// float64 => decimal numbers
	// bool => true or false
	// name := "John" | short way
	// %v is a placeholder for the variable
	// %d does the same but for decimal numbers
	// PRINT AND SCAN
	// Scan let the user input data
	// \n is a new line

	fmt.Printf("Enter\n your name: ")
	var name string
	fmt.Scan(&name)

	fmt.Printf("Hello %v, welcome to the game!", name)

	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)
}
