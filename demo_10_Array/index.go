package main

import (
	"fmt"
)

func main() {
	var array1 []int = []int{1, 2, 3, 4}
	fmt.Println(array1[0])

	var array2 []int = []int{5, 6, 7, 8}
	var array3 [3]string

	for index, item := range array1 {
		fmt.Println(index, item)
	}
	fmt.Println("----------------------------")
	for index, item := range array2 {

		fmt.Println(index, item)
	}
	fmt.Println("----------------------------")
	array3[0], array3[1], array3[2] = "android", "ios", "ipad"

	for index, item := range array3 {
		fmt.Println(index, item)
	}

}
