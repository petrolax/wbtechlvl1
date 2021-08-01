package main

import "fmt"

func main() {
	arr := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	num := make(chan int)
	numMul2 := make(chan int)

	go func() {
		for _, val := range arr {
			num <- val
			numMul2 <- val * 2
		}
		close(num)
		close(numMul2)
	}()

	for {
		res1, ok := <-num
		if !ok {
			break
		}
		res2 := <-numMul2
		fmt.Printf("x: %d; 2*x: %d\n", res1, res2)
	}
}
