//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	//* Store your favorite color in a variable using the `var` keyword
	var favoriteColor string = "green"
	fmt.Println("My favorite color is", favoriteColor)

	//* Store your birth year and age (in years) in two variables using
	//  compound assignment
	birthYear, age := 99, 23
	fmt.Println("My birth year is", birthYear, "My age is", age)

	//* Store your first & last initials in two variables using block assignment
	var (
		firstInitials string = "J"
		lastInitials  string = "N"
	)
	fmt.Println("initials =", firstInitials, lastInitials)

	//* Declare (but don't assign!) a variable for your age (in days),
	//  then assign it on the next line by multiplying 365 with the age
	var ageInday int
	ageInday = 365 * age
	fmt.Println("My age in day is", ageInday, "days")
	// 	variable created earlie

	fmt.Println()
}
