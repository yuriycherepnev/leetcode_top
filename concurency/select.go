//select с обработкой всех ответов

package main

import "fmt"

func double(numCh chan int) {
	value := <-numCh
	numCh <- value + value
}

func square(numCh chan int) {
	value := <-numCh
	numCh <- value * value
}

func cube(numCh chan int) {
	value := <-numCh
	numCh <- value * value * value
}

func main() {

	sqrCh := make(chan int)
	cubeCh := make(chan int)
	doubleCh := make(chan int)

	go square(sqrCh)
	go cube(cubeCh)
	go double(doubleCh)

	sqrCh <- 3
	cubeCh <- 5
	doubleCh <- 7

	for i := 1; i <= 3; i++ {
		select {

		case sqrVal := <-sqrCh:
			fmt.Println("Square:", sqrVal)

		case cubeVal := <-cubeCh:
			fmt.Println("Cube:", cubeVal)

		case doubleVal := <-doubleCh:
			fmt.Println("Double:", doubleVal)
		}
	}
}
