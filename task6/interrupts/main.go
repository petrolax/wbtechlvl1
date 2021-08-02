package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
)


func main() {
	ctx, cancel := context.WithCancel(context.Background())
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	channel := make(chan int)
	
	go func(c context.Context) {
		i := 0
		for {
			select {
			case channel <- i:
				i++
			case <-c.Done():
				fmt.Println("Goroutine was dead")
				close(channel)
				return
			}
		}
	}(ctx)
	
	go func() {
		for {
			select {
			case <-signals:
				cancel()
				return
			}
		}
	}()

	for val := range channel {
		fmt.Println(val)
	}
}