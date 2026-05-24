package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println(runtime.GOMAXPROCS(0))

	fmt.Println(runtime.NumCPU())

	runtime.GOMAXPROCS(16)

	fmt.Println(runtime.GOMAXPROCS(0))
	fmt.Println(runtime.NumCPU())

}
