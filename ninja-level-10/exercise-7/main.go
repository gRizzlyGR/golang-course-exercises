package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	numGoroutines := 10

	for i := 0; i < numGoroutines; i++ {
		go put(c, 10)
	}

	pull(c, 100)

	close(c)

	fmt.Println("about to exit")
}

func put(c chan<- int, limit int) {
	for i := 0; i < limit; i++ {
		c <- i
	}
}

func pull(c <-chan int, limit int) {
	for i := 0; i < limit; i++ {
		fmt.Println(i, <-c)
	}
}
