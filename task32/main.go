package main

import (
	"fmt"
	"time"
)

func sleep(sec int64) {
	tn := time.Now().Unix()
	fmt.Println(tn)
	for (time.Now().Unix() - tn) != sec {
	}
	fmt.Println(time.Now().Unix())
}

func main() {
	sleep(6)
}
