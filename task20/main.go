package main

import "fmt"

//Программа выведет [b b a] [a a]. Основная переменная slice не изменится, также, как и в случае с 18 заданием.
//После append создаётся новый срез, в который добавляется 3-е значение 'a', а первые 2 заменяются на значение 'b'
func main() {
	slice := []string{"a", "a"}

	func(slice []string) {
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice)
	}(slice)
	fmt.Print(slice)
}
