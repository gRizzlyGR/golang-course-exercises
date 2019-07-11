package main

import "fmt"

func main() {
	a := increase()
	b := increase()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}

func increase() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
