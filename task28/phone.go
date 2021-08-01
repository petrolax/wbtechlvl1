package main

import "fmt"

type AndroidPhone struct{}

func (ap AndroidPhone) InsertIntoMicroUSB() {
	fmt.Println("MicroUSB connector is plugged into Android Phone")
}
