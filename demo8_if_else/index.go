package main

import (
	"fmt"
)

func main() {
	//fnFor()
	fnSwitchCase()

}

func fnFor() {
	someValue := 10
	if someValue == 10 {
		fmt.Println("== 10")
	} else {
		fmt.Println("!==10")
	}

	if someValue > 10 || someValue < 20 {
		fmt.Println("someValue > 10 || someValue < 20 ")
	} else {
		fmt.Println("someValue < 10 || someValue > 20")
	}

	if result := returnValue(); result == 10 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}

func returnValue() int {
	return 10
}

func fnSwitchCase() {
	index := 0
	switch index {
	case 0:
		fmt.Println("0")
		break
	case 1:
		fmt.Println("1")
		break
	case 2:
		fmt.Println("2")
		break
	default:
		break
	}
}
