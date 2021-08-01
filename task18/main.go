package main

import "fmt"

// Данная функция поменяет 0 значение среза a, но после выполнения append создаётся новый срез с уже другим размером и указателем на массив.
// Это можно проверить, написав после append v[2] = 120, например, и основный срез также не изменится.
func someAction(v []int8, b int8) {
	v[0] = 100
	v = append(v, b)
}

func main() {
	var a = []int8{1, 2, 3, 4, 5}
	someAction(a, 6)
	fmt.Println(a)
}
