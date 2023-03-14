package main

import "fmt"

func main() {
	x := 100
	y := 200
	if x > y {
		fmt.Println("false")
	} else if x < y {
		fmt.Println("true")
	} else {
		fmt.Println("equal")
	}
}
