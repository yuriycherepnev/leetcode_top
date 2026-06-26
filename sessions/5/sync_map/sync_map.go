// Разработчик дал на ревью код своего нового кэша, нам необходимо провести код-ревью
// Кэш будет использоваться под высокой нагрузкой в проде
// Частота записи/чтения 20%/80% соответственно

package main

import (
	"fmt"
	"sync"
)

func main() {
	cacheMap := Cache{}
	fmt.Println(cacheMap.GetOrCreate("hello", "world"))
	fmt.Println(cacheMap.Get("hello"))
}

type Cache struct {
	cacheMap sync.Map
}

// GetOrCreate проверяет, существует ли ключ.
// Если он не существует, создается новое значение.
func (c *Cache) GetOrCreate(key, value string) string {
	store, _ := c.cacheMap.LoadOrStore(key, value)
	actual, ok := store.(string)
	if !ok {
		panic("cache value is not string")
	}
	return actual
}

func (c *Cache) Get(key string) string {
	v, _ := c.cacheMap.Load(key) // c.CacheM.Load(key)
	actual, ok := v.(string)
	if !ok {
		panic("cache value is not string")
	}
	return actual
}
