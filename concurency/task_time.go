// нужно вывести сколько секунд работала главная горутина mainSeconds
// и сколько суммарно секунд работала каждая отдельно totalWorkSeconds

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func randomWait() int {
	workSeconds := rand.Intn(10)

	time.Sleep(time.Duration(workSeconds) * time.Second)

	return workSeconds
}

// решение через каналы
/*
func main() {
	mainSeconds := 0
	totalWorkSeconds := 0
	start := time.Now()
	workChan := make(chan int)

	for range 100 {
		go func() {
			workSeconds := randomWait()

			workChan <- workSeconds
		}()
	}

	for range 100 {
		totalWorkSeconds += <-workChan
	}

	elapsed := time.Since(start)
	mainSeconds = int(elapsed.Seconds())

	fmt.Println(mainSeconds)
	fmt.Println(totalWorkSeconds)
}
*/

// решение через mutex и waitGroup

func main() {
	mainSeconds := 0
	totalWorkSeconds := 0
	mutex := sync.Mutex{}
	start := time.Now()
	wg := sync.WaitGroup{}

	wg.Add(100)
	for range 100 {
		go func() {
			defer wg.Done()
			workSeconds := randomWait()

			mutex.Lock()
			totalWorkSeconds += workSeconds
			mutex.Unlock()
		}()
	}

	wg.Wait()

	elapsed := time.Since(start)
	mainSeconds = int(elapsed.Seconds())

	fmt.Println(mainSeconds)
	fmt.Println(totalWorkSeconds)
}
