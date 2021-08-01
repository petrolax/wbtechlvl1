package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var someMap sync.Map

	for i := 0; i < 10; i++ {
		go func(idx int) {
			someMap.Store(idx, rand.Intn(100))
			fmt.Println("Worker", idx)
		}(i)
	}

	time.Sleep(time.Millisecond)

	someMap.Range(func(k, v interface{}) bool {
		fmt.Println("key:", k, ", val:", v)
		return true // if false, Range stops
	})

}
