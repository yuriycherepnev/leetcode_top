package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	groups := 7

	wg.Add(groups)
	for i := 1; i <= 7; i++ {
		go func(n int) {
			defer wg.Done()
			result := 0
			for j := 1; j <= n; j++ {
				result += j
			}
			fmt.Println(n, "-", result)
		}(i)
	}
	wg.Wait()
	fmt.Println("The End")
}

/*
package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	groups := 7

	wg.Add(groups)
	for i := 1; i <= groups; i++ {
		go sum(i, wg)
	}
	wg.Wait()

	fmt.Println("The End")
}

func sum(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	result := 0
	for i := 1; i <= n; i++ {
		result += i
	}
	fmt.Println(n, "-", result)
}
*/
