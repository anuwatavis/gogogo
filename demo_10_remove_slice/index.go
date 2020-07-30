package main

import "fmt"

func main() {
	var numbers = []int{1, 2, 3, 4, 5, 6}
	showSlice(numbers)

	// leading remove
	numbers = numbers[1:len(numbers)] //slice exclusive
	showSlice(numbers)

	numbers = numbers[1:len(numbers)] //slice exclusive
	showSlice(numbers)
	//tailing remove
	numbers = numbers[0 : len(numbers)-1]
	showSlice(numbers)
}

func showSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
