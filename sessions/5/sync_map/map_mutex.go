package main

import (
	"fmt"
	"sync"
)

func main() {
	cacheMap := CacheM{}
	fmt.Println(cacheMap.GetOrCreateM("hello", "world"))
	fmt.Println(cacheMap.GetM("hello"))
}

var cacheM = make(map[string]string)

type CacheM struct {
	CacheM sync.Map
	mutex  sync.RWMutex
}

func (c *CacheM) GetOrCreateM(key, value string) string {
	c.mutex.RLock()
	v, ok := cacheM[key]
	if ok {
		return v
	}
	c.mutex.RUnlock()

	c.mutex.Lock()
	defer c.mutex.Unlock()

	v, ok = cacheM[key]
	if ok {
		return v
	}

	cacheM[key] = value
	return value
}

func (c *CacheM) GetM(key string) string {
	c.mutex.RLock()
	defer c.mutex.Unlock()
	v, ok := cacheM[key]
	if ok {
		return v
	}
	return ""
}
