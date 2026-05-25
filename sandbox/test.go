package main

import (
	"fmt"
	"sync"
)

func main() {
	var cache sync.Map

	cache.Store("user:1", "Yuriy")

	value, ok := cache.Load("user:1")
	fmt.Println(value, ok)

}
