package main

import "fmt"

func main() {
	faveSport := "fifa"
	switch faveSport {
	case "basketbal":
		fmt.Println("indoors")
	case "hockey":
		fmt.Println("indoors")
	case "swimming":
		fmt.Println("both")
	case "running":
		fmt.Println("both")
	case "fifa":
		fmt.Println("both")
	}
}
