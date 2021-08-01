package main

import "fmt"

func printType(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("Type: int")
		break
	case string:
		fmt.Println("Type: string")
		break
	case chan int:
		fmt.Println("Type: chan int")
		break
	case bool:
		fmt.Println("Type: bool")
		break
	}
}

func main() {
	a, b, c, d := 0, "hello", make(chan int), true
	printType(c)
	printType(a)
	printType(d)
	printType(b)
}
