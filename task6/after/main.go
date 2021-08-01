package main

import (
	"fmt"
	"time"
)

func main() {
	nums := make(chan int)
	done := make(chan bool)
	go func() {
		i := 0
		for {
			select {
			case nums <- i:
				i++
			case <-done:
				fmt.Println("Goroutine was stop")
				close(nums)
				return
			}
		}
	}()

	go func() {
		select {
		case <-time.After(time.Second * 3):
			done <- true
		}
	}()

	for num := range nums {
		fmt.Println("Value from channel:", num)
	}
}
