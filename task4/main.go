package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
)

var (
	IsClosed = false
	once sync.Once
)

func worker(ctx context.Context, id int64, c chan string, wg *sync.WaitGroup) {
	for {
		select {
		case val, ok := <-c:
			if !ok {
				fmt.Println("Worker", id, "was dead")
				wg.Done()
				return
			}
			fmt.Printf("Worker%d: %s\n", id, val)
			fmt.Print("Enter data: ")
		case <-ctx.Done():
			once.Do(func() {
				close(c)
				IsClosed = true
			})
			
			fmt.Println("Worker", id, "was dead")
			wg.Done()
			return
		}
	}
}

func startWorkers(ctx context.Context, amount int64, c chan string, wg *sync.WaitGroup) {
	for i := int64(0); i < amount; i++ {
		wg.Add(1)
		go worker(ctx, i, c, wg)
	}
}

func main() {
	var amountOfWorkers int64
	fmt.Print("Enter amout of workers: ")
	fmt.Scanf("%d\n", &amountOfWorkers)
	if amountOfWorkers == 0 {
		panic("Workers can't be 0")
	}
	
	ctx, cancel := context.WithCancel(context.Background())
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	data := make(chan string)
	var wg sync.WaitGroup
	
	startWorkers(ctx, amountOfWorkers, data, &wg)
	go func() {
		for {
			select {
			case <-signals:
				cancel()
				return
			}
		}
	}()

	input := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter data: ")
	for input.Scan() {
		str := input.Text()
		if IsClosed { break}
		if str == "stop" || str == "STOP" {
			close(data)
			break
		}
		data <- str
	}
	wg.Wait()
}
