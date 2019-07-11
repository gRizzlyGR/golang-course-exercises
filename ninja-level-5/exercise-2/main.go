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

	m := map[string]person{
		person1.last: person1,
		person2.last: person2,
	}

	for key, p1 := range m {
		fmt.Printf("%v - %v %v likes: \n\t", key, p1.first, p1.last)
		for _, flavor := range p1.iceCreamFlavors {
			fmt.Printf("%v, ", flavor)
		}
		fmt.Println()
	}
}
