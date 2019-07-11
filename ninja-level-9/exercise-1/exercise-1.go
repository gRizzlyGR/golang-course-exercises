package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go func() {
		fmt.Println("First goroutine")
		wg.Done()
	}()

	go func() {
		fmt.Println("Second goroutine")
		wg.Done()
	}()

	fmt.Println("Goroutines:", runtime.NumGoroutine())
	wg.Wait()
}
