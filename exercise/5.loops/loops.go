//--Summary:
//  Implement the classic "FizzBuzz" problem using a `for` loop.
//
//--Requirements:
//* Print integers 1 to 50, except:
//  - Print "i Fizz" if the integer is divisible by 3
//  - Print "i Buzz" if the integer is divisible by 5
//  - Print "i FizzBuzz" if the integer is divisible by both 3 and 5
//
//--Notes:
//*
package main

import "fmt"

func main() {

	for i := 1; i <= 50; i++ {

		div3 := i%3 == 0
		div5 := i%5 == 0

		if div3 && div5 {

			fmt.Println(i, "Fizzbuzz")
		} else if div3 {
			fmt.Println(i, "Fizz")
		} else if div5 {
			fmt.Println(i, "Buzz")
		} else {
			fmt.Println(i)
		}

	}

}
