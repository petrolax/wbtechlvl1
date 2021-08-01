package main

import "fmt"

func checkUniq(src string) bool {
	for i := 0; i < len(src); i++ {
		for j := i + 1; j < len(src); j++ {
			if src[i] == src[j] {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(checkUniq("helo"))
}
