package main

import (
	"fmt"
	"time"
)

func main() {
	n := 0
	fmt.Print("Enter number of seconds for timeout: ")
	fmt.Scanf("%d", &n)
	nums := make(chan int)
	// timenow := time.Now().Unix()
	done := make(chan bool)

	go func() {
		i := 0
		for {
			select {
			case nums <- i:
				i++
			case <-done:
				fmt.Println("Timed out!")
				// fmt.Println("Seconds:", time.Now().Unix()-timenow)
				close(nums)
				return
			}
		}
	}()

	go func() {
		select {
		case <-time.After(time.Second * time.Duration(n)):
			done <- true
		}
	}()

	for num := range nums {
		fmt.Println("Value from channel:", num)
	}
}
