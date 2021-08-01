package main

import "fmt"

func removeIndexFromSlice(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	arr = removeIndexFromSlice(arr, 3)
	fmt.Println(arr)
}
