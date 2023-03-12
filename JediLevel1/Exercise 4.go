package main

import "fmt"

type hotdog int

var x hotdog

func main() {
	s := fmt.Sprintf("%v\n", x)
	fmt.Println(s)

	s = fmt.Sprintf("%T\n", x)
	fmt.Println(s)

	x = 42
	fmt.Println(x)
}
