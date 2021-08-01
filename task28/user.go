package main

import "fmt"

type User struct{}

func (u User) InsertUsbFlashDriverIntoDevice(dev Device) {
	fmt.Println("User insert USB flash driver into device")
	dev.InsertIntoUSB()
}
