package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 4, 6, 8, 10, 12}
	res := make(chan int)

	for _, i := range arr {
		go func(value int, c chan<- int) {
			res <- value * value
		}(i, res)
	}

	sum := 0
	count := 0
	for i := range res {
		sum += i

		count++
		if count == len(arr) {
			close(res)
		}
	}

	fmt.Println("Result of summ -", sum)
}
