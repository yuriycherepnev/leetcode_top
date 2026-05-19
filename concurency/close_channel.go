package main

import (
	"fmt"
	"sync"
)

func sendValue(wg *sync.WaitGroup, intChan chan int) {
	number := 5

	intChan <- number
	intChan <- number
	intChan <- number
	intChan <- number
	intChan <- number

	close(intChan)
	wg.Done()
}

func getValue(wg *sync.WaitGroup, intChan chan int) {
	// range автоматически читает значения до закрытия канала и выходит после
	//for value := range intChan {
	//	fmt.Println(value)
	//}

	// также можно обрабатывать сообщения из канала в бесконечном цикле
	//for {
	//	value, opened := <-intChan
	//	if !opened {
	//		break
	//	}
	//}

	//работа с буферизованным каналом наглядно
	for i := 0; i < 7; i++ {
		value, opened := <-intChan
		fmt.Println(value, opened)
	}

	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}
	intChan := make(chan int, 5)

	wg.Add(2)

	go sendValue(wg, intChan)
	go getValue(wg, intChan)

	wg.Wait()

}
