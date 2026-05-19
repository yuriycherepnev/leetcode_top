package main

import (
	"fmt"
	"sync"
)

var counter2 int = 0
var wg sync.WaitGroup

func work(number int, mutex *sync.Mutex) {
	mutex.Lock()
	counter2 = 0
	for k := 1; k <= 5; k++ {
		counter2 += 1
		fmt.Println("Goroutine", number, "-", counter2)
	}
	mutex.Unlock()
	wg.Done()
}

func main() {
	var mutex sync.Mutex
	goroutinesCount := 4
	wg.Add(goroutinesCount)
	for i := 1; i <= goroutinesCount; i++ {
		go work(i, &mutex)
	}
	wg.Wait()
}
