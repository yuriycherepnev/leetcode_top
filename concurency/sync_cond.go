package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var pokemonList = []string{"Charmander", "Squirtle", "Bulbasaur", "Jigglypuff"}
var cond = sync.NewCond(&sync.Mutex{})
var pokemon = ""

func main() {
	go func() {
		cond.L.Lock()
		defer cond.L.Unlock()

		fmt.Println(pokemon)
		for pokemon != "Pikachu" {
			cond.Wait()
		}
		println("Caught" + pokemon)
		pokemon = ""
	}()

	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(time.Millisecond)

			cond.L.Lock()
			pokemon = pokemonList[rand.Intn(len(pokemonList))]
			cond.L.Unlock()

			cond.Signal()
		}
	}()

	time.Sleep(100 * time.Millisecond) // lazy wait
}
