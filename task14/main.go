package main

import "fmt"

func toSet(arr []string) []string {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				arr = append(arr[:j], arr[j+1:]...)
				j--
			}
		}
	}
	return arr
}

func main() {
	arr := []string{"cat", "cat", "dog", "value", "cat", "tree", "tree", "dog", "value"}
	res := toSet(arr)
	fmt.Println(res)
}
