package main

import (
	"fmt"
	"sync"
)

func simpleGo(wg *sync.WaitGroup) {
	fmt.Println(4)
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}
	//wg.Add()

	go simpleGo(wg)
	wg.Wait()
}
