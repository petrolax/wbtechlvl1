package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	value int32
}

func main() {
	counter := new(Counter)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int, wg *sync.WaitGroup) {
			atomic.AddInt32(&counter.value, 1)
			fmt.Println("id:", id, "value:", atomic.LoadInt32(&counter.value))
			wg.Done()
		}(i, &wg)
	}

	wg.Wait()
}
