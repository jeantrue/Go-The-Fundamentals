//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//  - Products must include the price and the name
//* Insert 3 products into the array

package main

import "fmt"

type Product struct {
	price int
	name  string
}

//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
func printStats(list [4]Product) {
	var cost, totalItem int

	for i := 0; i < len(list); i++ {
		item := list[i]
		cost += item.price

		if item.name != "" {
			totalItem += 1
		}

		fmt.Println("Last time on the list", list[totalItem-1])
		fmt.Println("Total item", totalItem)
		fmt.Println("Total cost", cost)
	}
}

func main() {

	//* Using an array, create a shopping list with enough room
	//  for 4 products
	shoppingList := [4]Product{
		{1, "book"},
		{6, "pen"},
		{5, "monitor"},
	}

	printStats(shoppingList)
	//* Add a fourth product to the list and print out the
	//  information again

	shoppingList[3] = Product{4, "I dont know"}
	printStats(shoppingList)

}
