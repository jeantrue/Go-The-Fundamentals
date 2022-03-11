package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = keyboard.Close()

	}()

	coffees := make(map[int]string)
	coffees[1] = "Cappucino"
	coffees[2] = "Latte"
	coffees[3] = "Americano"
	coffees[4] = "Mocha"
	coffees[5] = "Macchiato"
	coffees[6] = "Espresso"

	fmt.Println("MENU")
	fmt.Println("----")
	fmt.Println("1 - Cappucino")
	fmt.Println("2 - Latte")
	fmt.Println("3 - Americano")
	fmt.Println("4 - Mocha")
	fmt.Println("5 - Macchiato")
	fmt.Println("6 - Espresso")
	fmt.Println("Q - Quit the program")

	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if char == 'q' || char == 'Q' {
			break
		}

		i, _ := strconv.Atoi(string(char))
		t := fmt.Sprintf("You chose %s", coffees[i])
		fmt.Println(t)

	}

	fmt.Println("Program exiting...")
}

// func main() {
// 	reader := bufio.NewReader(os.Stdin)

// 	for {
// 		fmt.Print("-> ")

// 		// \n which is the enter key
// 		// _ mean ignor the second thing being return by that func, just discard it
// 		userInput, _ := reader.ReadString('\r')
// 		// going to replace user input is our source string
// 		//looking for  "\r" (enter for window (\n fot mac) ) to repalce it with "" (notthing)
// 		// -1 mean doing it everywhere
// 		userInput = strings.Replace(userInput, "\r", "", -1)

// 		if userInput == "quit" {
// 			break
// 		} else {
// 			fmt.Println(userInput)

// 		}
// 	}
// }
