package main

import "fmt"

type PhoneAdapter struct {
	androidp *AndroidPhone
}

func (pa PhoneAdapter) InsertIntoUSB() {
	fmt.Println("Adapter converts USB signal to microUSB")
	pa.androidp.InsertIntoMicroUSB()
}
