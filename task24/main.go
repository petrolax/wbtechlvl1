package main

import "fmt"

// Task: Создать слайс с предварительно выделенными 100 элементами.

func main() {
	slice := make([]int, 100)
	fmt.Println(len(slice), cap(slice))
}
