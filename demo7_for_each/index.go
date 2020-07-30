package main

import (
	"fmt"
)

func main() {
	course := []string{"Android", "IOS", "React"}
	for index, item := range course {
		fmt.Printf("%d %s \n", index, item)
	}

	for _, item := range course {
		fmt.Printf("%s \n", item)
	}
	fmt.Println("-----------------")
	fnForUsingBreak()
}

func fnForUsingBreak() {
	index := 0
	for true {
		if index > 5 {
			break
		}
		fmt.Printf("Index %d \n", index)
		index++
	}
}
