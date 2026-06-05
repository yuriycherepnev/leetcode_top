package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	count := atomic.Int32{}

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count.Add(1)
		}()
	}

	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count.Add(-1)
		}()
	}

	wg.Wait()

	result := count.Load()
	fmt.Println(result)
}
