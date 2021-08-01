package main

import (
	"fmt"
	"sort"
)

func binarySearch(arr []int, value int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := high - low/2
		guess := arr[mid]

		if guess == value {
			return mid
		}
		if guess > value {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	arr := []int{3, 14, 21, 532, 62, 4, 41, 32, 1, 3, 25, 3, 61, 3213}
	sort.Ints(arr)
	fmt.Println(arr)
	fmt.Println(binarySearch(arr, 25))
}
