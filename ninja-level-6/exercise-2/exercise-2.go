package main

import (
	"fmt"
)

func main() {
	fmt.Println(foo(1, 2, 3, 4, 5))

	x := []int{6, 7, 8, 9, 10}

	fmt.Println(foo(x...))
	fmt.Println(bar(x))
}

func foo(s ...int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}

	return sum
}

func bar(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}

	return sum
}
