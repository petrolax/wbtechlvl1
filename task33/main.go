package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randNums := make(chan int)
	evenNums := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			randNums <- rand.Intn(1000000)
		}
		close(randNums)
	}()

	go func() {
		for {
			val, ok := <-randNums
			if !ok {
				close(evenNums)
				return
			}
			if val%2 == 0 {
				evenNums <- val
			}
		}
	}()

	for value := range evenNums {
		fmt.Println("From even channel:", value)
	}
}
