package main

import (
	"fmt"
)

type person struct {
	first           string
	last            string
	iceCreamFlavors []string
}

func main() {
	person1 := person{
		first:           "Giuseppe",
		last:            "Rizzi",
		iceCreamFlavors: []string{"coffee", "hazelnut", "wild berries"},
	}

	person2 := person{
		first:           "Mario",
		last:            "Rossi",
		iceCreamFlavors: []string{"strawberry", "lemon", "pistachio"},
	}

	fmt.Printf("%v %v likes: \n\t", person1.first, person1.last)
	for _, flavor := range person1.iceCreamFlavors {
		fmt.Printf("%v, ", flavor)
	}

	fmt.Println()

	fmt.Printf("%v %v likes: \n\t", person2.first, person2.last)
	for _, flavor := range person2.iceCreamFlavors {
		fmt.Printf("%v, ", flavor)
	}

	fmt.Println()
}
