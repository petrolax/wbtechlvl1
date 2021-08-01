package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	} else if len(arr) == 2 {
		if arr[0] > arr[1] {
			arr[0], arr[1] = arr[1], arr[0]
		}
		return arr
	}

	arrLess, arrGreater := make([]int, 0), make([]int, 0)
	ref := arr[0]
	for i := 1; i < len(arr); i++ {
		if ref > arr[i] {
			arrLess = append(arrLess, arr[i])
		} else {
			arrGreater = append(arrGreater, arr[i])
		}
	}

	less := quickSort(arrLess)
	greater := quickSort(arrGreater)
	less = append(less, ref)
	return append(less, greater...)

}

func main() {
	arr := []int{67, 12, 4, 23, 412, 35, 6, 132, 151, 6, 123, 1, 534, 213}
	fmt.Println(quickSort(arr))
}
