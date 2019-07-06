package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func main() {
	me := person{
		first: "Giuseppe",
		last:  "Rizzi",
		age:   31,
	}

	fmt.Println(me)
	changeMe(&me, "Beppe")
	fmt.Println(me)
}

func changeMe(p *person, newFirst string) {
	p.first = newFirst
}
