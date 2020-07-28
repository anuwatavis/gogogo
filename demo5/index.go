package main

import (
	"fmt"
)

func main() {
	fmt.Println("Anuwataravis")
	display()
	fn2("Good Morning")
	fn3("number is", 2)
}

func display() {
	fmt.Println("display fucntion")

}

func fn2(msg string) {
	fmt.Println(msg)
}

func fn3(title string, version int) {
	fmt.Println(title)
	fmt.Println(version)
}
