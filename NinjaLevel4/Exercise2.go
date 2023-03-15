package main

import "fmt"

func main() {

	// slice
	x := []int{42, 43, 44, 45, 46, 47, 47, 48, 49, 50}

	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}
