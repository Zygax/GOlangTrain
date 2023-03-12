package main

// Raw strings are enclosed in back-ticks `. Here, \t and \n has no special meaning, they are considered as backslash with t and backslash with n. If you need to include backslashes, double quotes or newlines in your string, use a raw string literal.
func main() {
	a := `Raw strings are enclosed in back-ticks. 
	"Here, \t and \n has no special meaning, they are considered 
	as backslash with t and backslash with n." If you need to include backslashes, 
	double quotes or newlines in your string, use a raw string literal.`

	println(a)
}
