package main

import (
	"fmt"
)

func main() {
	foo(func() int {
		return 42
	})
}

func foo(f func() int) {
	x := f()
	x++
	fmt.Println(x)
}
