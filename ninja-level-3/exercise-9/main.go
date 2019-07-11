package main

import (
	"fmt"
)

func main() {
	favSport := "beach volley"

	switch favSport {
	case "soccer":
		fmt.Println("I like soccer")
	case "basket":
		fmt.Println("I like basket")
	default:
		fmt.Println("Neither, I like " + favSport)
	}

}
