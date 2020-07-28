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
	fmt.Printf("x = %d and y = %d | a = %d b = %d", x, y, a, b)

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
