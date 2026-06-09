package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	ready := false
	go func() {
		mu.Lock()
		for !ready {
			fmt.Println(111)
			cond.Wait()
		}
		fmt.Println("Условие выполнено, продолжаем работу.")
		mu.Unlock()
	}()

	time.Sleep(1 * time.Second)

	mu.Lock()
	//ready = true
	cond.Signal()
	mu.Unlock()

	time.Sleep(1 * time.Second)
}
