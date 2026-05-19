package main

import (
	"fmt"
	"sync"
)

func getVal(wg *sync.WaitGroup, getCh <-chan int) {
	fmt.Println(<-getCh)
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}

	getCh := make(<-chan int)
	wg.Add(1)
	go getVal(wg, getCh)
	wg.Wait()
}
