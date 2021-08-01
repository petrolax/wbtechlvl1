package main

import (
	"fmt"
	"sync"
)

//Данная программа закончится deadlock'ом, т.к. в горутину передаётся копия переменной wg и wg.Done() выполняется для копии.
//Чтобы исправить это нам необходимо чтобы функция принимала указатель на переменную
func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}
