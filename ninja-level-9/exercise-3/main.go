package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0
	gs := 100

	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < gs; i++ {

		go func() {
			v := counter
			fmt.Println(counter)
			runtime.Gosched()

			v++
			counter = v // race condition

			wg.Done()
		}()

		//fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Counter:", counter)

}
