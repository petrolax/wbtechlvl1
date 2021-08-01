package main

import (
	"fmt"
)

func reverseString(src string) (res string) {
	for _, symb := range src {
		res = string(symb) + res
	}
	return res
}

func main() {
	fmt.Println("Hello, 世界")
	fmt.Println(reverseString("Hello, 世界"))
	fmt.Println()
	fmt.Println("你好好好")
	fmt.Println(reverseString("你好好好"))
	fmt.Println()
	fmt.Println("\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98")
	fmt.Println(reverseString("\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"))
}
