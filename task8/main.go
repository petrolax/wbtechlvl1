package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	var nbit uint8
	fmt.Scanf("%d", &nbit)

	value := rand.Int63n(100)
	fmt.Println(value | rand.Int63n(2)<<nbit)
	fmt.Println(math.Pow(2, 32))
}
