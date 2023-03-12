package main

import "fmt"

func main() {

	a := 42 == 42
	b := 42 <= 43
	c := 42 >= 32
	d := 42 < 45
	e := 42 > 33
	f := 42 != 42

	fmt.Println(a, b, c, d, e, f)
}
