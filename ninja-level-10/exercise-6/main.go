package main

import "fmt"

func main() {
	c := make(chan int)

	go put(c, 100)

	pull(c)

	fmt.Println("about to exit")
}

func put(c chan<- int, limit int) {
	for i := 0; i < limit; i++ {
		c <- i
	}
	close(c)
}

func pull(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
