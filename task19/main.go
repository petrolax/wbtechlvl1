package main

var justString string

/*
*	1) Нужно объявить и описать функцию createHugeString, т.к. программа будет ругаться, что мы пытаемся вызвать функцию, которой нет
*	2) Не факт что строка будет иметь 100 символов (сработает паника, если значений меньше 100, т.к. мы пытаемся выйти за границы строки)
 */
func someFunc() {
	v := "Hello beautiful world!" //createHugeString(1 << 10)
	justString = v[:15]
}

func main() {
	someFunc()
}
