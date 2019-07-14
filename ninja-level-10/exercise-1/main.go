package main

import "fmt"

func main() {
	withGoroutine() // preferred approach

	withBuffer()
}

func withGoroutine() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

func withBuffer() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
