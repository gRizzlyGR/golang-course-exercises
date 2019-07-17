package main

import "fmt"

type customError struct {
	s string
}

func (e customError) Error() string {
	return e.s
}

func foo(e error) {
	fmt.Println(e)
}

func main() {
	ce := customError{
		s: "I am a custom error",
	}

	foo(ce)
}
