package main

import (
	"fmt"

	"point"
)

func main() {
	p1 := point.CreatePoint(1, 2)
	p2 := point.CreatePoint(4, 3)
	res := p1.GetDistance(*p2)
	fmt.Println(res)
}
