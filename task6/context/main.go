package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	nums := make(chan int)

	go func(ctx context.Context) {
		i := 0
		for {
			select {
			case nums <- i:
				i++
			case <-ctx.Done():
				fmt.Println("Goroutine was stop")
				close(nums)
				return
			}
		}
	}(ctx)

	go func() {
		time.Sleep(time.Second * 3)
		cancel()
	}()

	for num := range nums {
		fmt.Println("Value from channel:", num)
	}
}
