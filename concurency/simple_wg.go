package main

import (
	"fmt"
	"sync"
	"time"
)

func one(wg *sync.WaitGroup) {
	wg.Wait()
	wg.Add(1)
	fmt.Println(1)
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}
	go one(wg)
	go one(wg)
	go one(wg)

	time.Sleep(1 * time.Second)
}
