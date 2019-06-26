package main

import (
	"fmt"
)

const (
	a int = 42
	b     = 42.1
)

func main() {
	fmt.Printf("%v of type %T\n", a, a)
	fmt.Printf("%v of type %T\n", b, b)
}
