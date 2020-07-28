package main

import (
	"fmt"
)

func main() {

	//Explicit Declaration
	var tmp1 int = 0
	var tmp2 string = "hello"
	var tmp3 bool = true

	//const tem4 = 0

	//Implicit Declaration

	temp5 := 0

	fmt.Println(tmp1)
	fmt.Println(tmp2)
	fmt.Println(tmp3)
	fmt.Println(temp5)
	run()

}

var count int = 0

func run() {
	count++
	fmt.Println(count)
}
