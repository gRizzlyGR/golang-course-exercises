package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	numGoroutines := 10
	limit := 10

	putWithCallback(c, numGoroutines, limit)

	pullWithRange(c)

	fmt.Println("About to exit")
}

func putWithCallback(c chan<- int, numGoroutines, limit int) {
	done := 0
	iAmDone := func() {
		done++
		fmt.Println("Num done:", done)

		if done == numGoroutines {
			fmt.Println("Closing channel")
			close(c)
			return
		}
	}

	put := func(c chan<- int, limit int, callback func()) {
		for i := 0; i < limit; i++ {
			c <- i
		}
		callback()
	}

	for i := 0; i < numGoroutines; i++ {
		go put(c, limit, iAmDone)
	}

}

func pullWithRange(c <-chan int) {
	i := 0
	for v := range c {
		fmt.Println(i, v)
		i++
	}
}
