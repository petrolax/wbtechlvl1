package main

import (
	"fmt"
	"strings"
)

func reverseWordsAtStr(src string) (res string) {
	words := strings.Split(src, " ")
	for i := len(words) - 1; i > -1; i-- {
		if i == 0 {
			res += words[i]
			break
		}
		res += words[i] + " "
	}
	return res
}

func main() {
	exampleString := "snow dog sun"
	fmt.Println(reverseWordsAtStr(exampleString))
}
