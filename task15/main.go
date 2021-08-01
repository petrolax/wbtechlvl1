package main

import "fmt"

func main() {
	a, b := 10, 15
	a, b = b, a
	fmt.Println("a:", a, "b:", b)
}
