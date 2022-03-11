package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
)

var reader *bufio.Reader

type User struct {
	UserName        string
	Age             int
	FavouriteNumber float64
	OwnADog         bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	var user User

	user.UserName = readeString("What is your name ?")
	user.Age = readInt("How old are you ?")
	user.FavouriteNumber = readFloat("What is your favourite number ?")
	user.OwnADog = readBool("Do you own a dog ?")

	fmt.Println(fmt.Sprintf("Your name is %s. You are %d years old. Your favourite number is %.2f, Owns a dog: %t",
		user.UserName,
		user.Age,
		user.FavouriteNumber,
		user.OwnADog))

}

func readeString(s string) string {

	for {

		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		if userInput == "" {
			fmt.Println("Please enter a value")
		} else {
			return userInput
		}
	}
}

func prompt() {
	fmt.Print("-> ")
}

func readInt(s string) int {

	for {

		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		//string to Int
		num, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Please enter a whole number")
		} else {
			return num
		}

	}
}

func readFloat(s string) float64 {

	for {

		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		//string to Int
		num, err := strconv.ParseFloat(userInput, 64)
		if err != nil {
			fmt.Println("Please enter a number")
		} else {
			return num
		}

	}
}

func readBool(s string) bool {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		fmt.Println(s)
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)

		}
		if strings.ToLower(string(char)) != "y" && strings.ToLower(string(char)) != "n" {
			fmt.Print("Please type y or n. ")
		} else if char == 'n' || char == 'N' {
			return false
		} else if char == 'y' || char == 'Y' {
			return true
		}
	}
}
