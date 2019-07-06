package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("Hi, I'm %v %v and I'm %v old.\n", p.first, p.last, p.age)
}

func main() {
	me := person{
		first: "Giuseppe",
		last:  "Rizzi",
		age:   31,
	}

	me.speak()
}
