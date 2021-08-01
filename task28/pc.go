package main

import "fmt"

type PC struct{}

func (pc PC) InsertIntoUSB() {
	fmt.Println("USB connector is plugged into PC")
}
