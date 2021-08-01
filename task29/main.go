package main

import (
	"fmt"
	"math/big"
)

func main() {
	a, b := big.NewInt(1<<32), big.NewInt(1<<33)
	fmt.Println(a.Add(a, b))
	fmt.Println(a.Sub(a, b))
	fmt.Println(a.Mul(a, b))
	fmt.Println(a.Div(a, b))
}
