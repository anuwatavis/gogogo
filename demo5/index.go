package main

import (
	"fmt"
)

func main() {
	fmt.Println("Anuwataravis")
	display()
	fn2("Good Morning")
	fn3("number is", 2)
	//return function
	const a, b = 2, 3
	//d stand for decimal
	fmt.Printf("%d+%d=%d\n", a, b, sum(a, b))
	var x, y int = swap(a, b)
	fmt.Printf("x = %d and y = %d | a = %d b = %d \n", x, y, a, b)
	_x, _y, title := swap2(10, 20)
	fmt.Printf("x = %d and y = %d, string = %s | a = %d b = %d", _x, _y, title, 10, 20)

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

func sum(a int, b int) int {
	return a + b
}

func swap(a int, b int) (int, int) {
	return b, a
}

func swap2(a int, b int) (x, y int, title string) {
	y = a
	x = b
	title = "Result"
	return
}
