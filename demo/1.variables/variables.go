package main

import "fmt"

func main() {
	var myName = "Jean"
	fmt.Println("My name is", myName, myName)

	var Gname string = "Nara"
	fmt.Println("name =", Gname)

	userName := "admin"
	fmt.Println("username is =", userName)

	var sum int
	fmt.Println("The sum is", sum)

	manyvar, manyvar2 := 1, 2
	fmt.Println("manyvar_1 is", manyvar, "and manyvar_2 is", manyvar2)

	aonther, manyvar2 := 2, 0
	fmt.Println("other is", aonther, "and mayvar_2 that reassign is", manyvar2)

	var (
		bookName = "Variable"
		bookType = "I dont know"
	)

	fmt.Println("Book Name is", bookName)
	fmt.Println("Book Type is", bookType)

	word1, word2, _ := "hello", "world", "!"
	fmt.Println(word1, word2)
}
