// Разработчик дал на ревью код своего нового кэша, нам необходимо провести код-ревью
// Кэш будет использоваться под высокой нагрузкой в проде
// Частота записи/чтения 20%/80% соответственно

package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println(GetOrCreate("hello", "world"))
	fmt.Println(Get("hello"))
}

var cache = make(map[string]string)

type Cache struct {
	CacheM sync.Map
	sync.RWMutex
}

// GetOrCreate checks whether the key exists.
// If it does not exist, it creates a new value.
func (с *Cache) GetOrCreate(key, value string) string {

	m.Lock()
	value = cache[key]
	m.Unlock()

	if value != "" {
		return value
	}

	m.Lock()
	cache[key] = value
	m.Unlock()

	return value
}

func (с *Cache) Get(key string) string {

	v := c.CacheM[key] // c.CacheM.Load(key)
	return v
}
