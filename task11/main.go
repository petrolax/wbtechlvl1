package main

import "fmt"

func Intersection(a, b []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				res = append(res, a[i])
			}
		}
	}
	return res
}

func main() {
	a := []int{2, 4, 3, 5, 6, 7}
	b := []int{9, 2, 7, 6}

	fmt.Println(Intersection(a, b))
}
