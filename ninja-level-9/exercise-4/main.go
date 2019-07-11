package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 0
	gs := 100

	var wg sync.WaitGroup
	wg.Add(100)

	var mtx sync.Mutex

	for i := 0; i < gs; i++ {

		go func() {
			mtx.Lock()

			v := counter
			fmt.Println(counter)
			// runtime.Gosched() it does not make sense since we're locking

			v++
			counter = v // mutex prevents race condition

			mtx.Unlock()

			wg.Done()
		}()

		//fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Counter:", counter)

}
