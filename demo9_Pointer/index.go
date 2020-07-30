package main

import (
	"fmt"
)

func main() {
	message := "some message"
	var msgPointer *string = &message //& point to address of message variable
	fmt.Println(message)
	fmt.Println(*msgPointer) // *  use for get value from address of pointer

	changeMessage(msgPointer)
	fmt.Println(message)
}

func changeMessage(aPointer *string) {
	*aPointer = "New message"
}
