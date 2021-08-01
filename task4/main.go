package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func worker(id int, c <-chan string, wg *sync.WaitGroup) {
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
		}
	}
}

func startWorkers(amount int, c <-chan string, wg *sync.WaitGroup) {
	for i := 0; i < amount; i++ {
		wg.Add(1)
		go worker(i, c, wg)
	}
}

func main() {
	data := make(chan string)
	var wg sync.WaitGroup
	amountOfWorkers := 0
	fmt.Print("Enter amout of workers: ")
	fmt.Scanf("%d\n", &amountOfWorkers)
	if amountOfWorkers == 0 {
		panic("Workers can't be 0")
	}

	startWorkers(amountOfWorkers, data, &wg)

	input := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter data: ")
	for input.Scan() {
		str := input.Text()
		if str == "stop" || str == "STOP" {
			close(data)
			break
		}
		data <- str
		// fmt.Print("Enter data: ")
	}
	wg.Wait()
}
