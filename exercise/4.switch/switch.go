//--Summary:
//  Create a program to display a classification based on age.
//
//--Requirements:
//* Use a `switch` statement to print the following:
//  - "newborn" when age is 0
//  - "toddler" when age is 1, 2, or 3
//  - "child" when age is 4 through 12
//  - "teenager" when age is 13 through 17
//  - "adult" when age is 18+

package main

import "fmt"

func age() int {
	return 19
}
func main() {
	//* Use a `switch` statement to print the following:
	switch yearAge := age(); {
	case yearAge == 0:
		//  - "newborn" when age is 0
		fmt.Println("newborn")
	case yearAge >= 1 && yearAge <= 3:
		//  - "toddler" when age is 1, 2, or 3
		fmt.Println("toddler")
	case yearAge < 13:
		//  - "child" when age is 4 through 12
		fmt.Println("child")
	case yearAge < 18:
		//  - "teenager" when age is 13 through 17
		fmt.Println("teenager")
	default:
		//  - "adult" when age is 18+
		fmt.Println("adult")
	}
}
