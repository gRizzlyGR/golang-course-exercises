package main

import (
	"fmt"
)

func main() {
	switch {
	case 1 == 2:
		fmt.Println("true 1")
	case 2 == 3:
		fmt.Println("true 2")
	case 3 == 3:
		fmt.Println("true 3")
	}
}
