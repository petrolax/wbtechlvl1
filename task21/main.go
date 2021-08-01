package main

import (
	"fmt"
	"sync"
)

//	Task: Написать программу, которая в конкурентном виде читает элементы из массива в stdout.

func main() {
	arr := [10]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	var wg sync.WaitGroup

	for i := 0; i < len(arr); i++ {
		wg.Add(1)
		go func(id int, wg *sync.WaitGroup) {
			for _, val := range arr {
				fmt.Println("Id:", id, "value:", val)
			}
			wg.Done()
		}(i, &wg)
	}
	wg.Wait()
}
