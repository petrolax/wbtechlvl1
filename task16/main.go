package main

import "fmt"

//Данная программа выведет 0, т.к. внутри условного оператора if создаётся локальная переменная n со значением 1 и инкрементируется. За пределами оператора данной переменной не существует.
func main() {
	n := 0
	if true {
		n := 1
		n++
	}
	fmt.Println(n)
}
