package main

import "fmt"

// TYPED constants. A const declared specifying the type in the declaration is a typed constant.
const a int32 = 8
const b string = "SD energy"

// UNTYPED constants. An untyped constant is a constant whose type has not been specified.
const c = 3.1423243242
const d = "who let the dogs out"

func main() {
	fmt.Println(a, b, c, d)
}
