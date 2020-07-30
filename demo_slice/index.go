package main

import "fmt"

func main() {

	//	var numbers = make([]int, 0, 5) // 3 len , 5 cap // if out of index , auto to increase with size
	// : 5 + 5 =  10 if out of index from default that is 5
	var numbers []int
	numbers = append(numbers, 1)
	numbers = append(numbers, 2)
	numbers = append(numbers, 2, 3)
	showSlice(numbers)

}

func showSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
