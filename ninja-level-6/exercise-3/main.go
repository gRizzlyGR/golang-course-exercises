package main

import (
	"fmt"
)

func main() {
	defer foo()

	fmt.Println("I'm first!")
}

func foo() {
	fmt.Println("I'm last!")
}
