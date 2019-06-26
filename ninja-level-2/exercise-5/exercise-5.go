package main

import (
	"fmt"
)

func main() {
	a := `raw string
	
	everything is verbatim
		here
		"even quotes"
	`
	fmt.Println(a)
}
