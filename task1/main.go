package main

import "fmt"

type Human struct {
	name string
}

type Action struct {
	Human
}

func (a Action) SayHello() {
	if a.name == "" {
		a.name = "Human"
	}
	fmt.Printf("%s say 'Hello'\n", a.name)
}

func main() {
	human := Human{
		name: "Anya",
	}

	action := Action{
		human,
	}

	action.SayHello()

	a1 := Action{}
	a1.SayHello()
}
