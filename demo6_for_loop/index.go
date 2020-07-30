package main

import (
	"fmt"
)

func main() {
	fnWhile()

}

func fnFor() {
	for index := 0; index <= 10; index++ {
		fmt.Printf("index%d \n", index)
	}
}

func fnWhile() {
	index := 0
	for index < 5 {
		index++
		fmt.Printf("Index %d \n", index)
	}

}
