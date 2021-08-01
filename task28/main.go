package main

import "fmt"

func main() {
	user := User{}

	pc := PC{}
	aPhone := AndroidPhone{}
	aPhoneAdapter := PhoneAdapter{&aPhone}

	user.InsertUsbFlashDriverIntoDevice(pc)
	fmt.Println()
	user.InsertUsbFlashDriverIntoDevice(aPhoneAdapter)
}
