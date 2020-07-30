package main

import "fmt"

func main() {
	var numbers = map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println("", numbers["two"])

	//dynamic map

	var colors = make(map[string]string)
	colors["red"] = "#F00"
	colors["green"] = "#0F0"
	colors["blue"] = "#00F"
	fmt.Println("", colors)
	fmt.Println("", colors["green"])

	//multiple map
	var course = make(map[string]map[string]int)
	course["Android"] = make(map[string]int)
	course["Android"]["price"] = 200
	course["Android"]["height"] = 30

	fmt.Println("course -->", course)
}
