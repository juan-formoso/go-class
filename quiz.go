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

		LOGICAL OPERATORS
		&& is and
		|| is or
		! is not
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

	score := 0
	num_questions := 2

	fmt.Printf("What is better, the RTX 3080 or RTX 3090? ")
	var answer string
	var answer2 string // go will not allow you to have the same variable name in different scopes
	fmt.Scan(&answer, &answer2)

	// We need to use " " because it has a space between the words and normally go does not recognize it
	if answer+" "+answer2 == "RTX 3090" || answer+" "+answer2 == "rtx 3090" {
		fmt.Println("You are correct!")
		score += 1 // we can also use score++ (but it works as an implement) or score = score + 1
	} else {
		fmt.Println("You are wrong!")
	}

	fmt.Printf("How many cores does the Ryzen 9 3900X have? ")
	var cores uint
	fmt.Scan(&cores)

	if cores == 12 {
		fmt.Println("You are correct!")
		score++
	} else {
		fmt.Println("You are wrong!")
	}

	fmt.Printf("You scored %v our of %v.\n", score, num_questions)
	percent := (float64(score) / float64(num_questions)) * 100
	fmt.Printf("That is %v%%. percent!", percent)
}
