package main

import "myapps/packageone"

var myVar string = "myVar-variable is package level from main "

func main() {
	var blockVar string = "This is a block level variable"

	//this is a export function form packageone
	packageone.Printme(myVar, blockVar)

}
