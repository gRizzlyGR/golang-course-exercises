package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)
	numGoroutines := 10

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go put(c, 10, &wg)
	}

	go pull(c)

	wg.Wait()
	close(c)

	fmt.Println("about to exit")
}

func put(c chan<- int, limit int, wg *sync.WaitGroup) {
	for i := 0; i < limit; i++ {
		c <- i
	}
	wg.Done()
}

func pull(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
