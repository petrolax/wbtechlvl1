package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for i := 0; i < len(arr); i++ {
		wg.Add(1)
		go func(value int) {
			defer wg.Done()
			fmt.Printf("%d power of 2: %d\n", value, value*value)
		}(arr[i])
	}
	wg.Wait()
}
