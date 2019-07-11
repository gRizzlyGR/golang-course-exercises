package main

import (
	"fmt"
)

func main() {
	x := foo

	y := func(name string) {
		fmt.Println("Hi", name)
	}

	x()
	y("Mark")
}

func foo() {
	fmt.Println("I can be assigned!")
}
