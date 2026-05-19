package main

import (
	"fmt"
	"sync"
)

var counter int

func increment(mutex *sync.RWMutex, wg *sync.WaitGroup) {
	for i := 0; i < 100000; i++ {
		mutex.RLock()
		counter++
		mutex.RUnlock()
	}
	wg.Done()
}

func main() {
	mutex := &sync.RWMutex{}
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go increment(mutex, wg)
	go increment(mutex, wg)
	wg.Wait()

	fmt.Println(counter)
}
