package main

import "fmt"

type Room struct {
	name    string
	cleaned bool
}

func checkCleanliness(rooms [4]Room) {

	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		if room.cleaned {
			fmt.Println(room.name, "is cleaned")
		} else {
			fmt.Println(room.name, "is dirty")
		}

	}
	fmt.Println("")
}
func main() {

	allOfroom := [...]Room{
		{name: "Bed Room"},
		{name: "Office Room"},
		{name: "Reception Room"},
		{name: "Lobby Room"},
	}

	checkCleanliness(allOfroom)
	fmt.Println("Performing cleaning..")
	allOfroom[0].cleaned = true
	allOfroom[2].cleaned = true
	allOfroom[3].cleaned = true
	checkCleanliness(allOfroom)
}
