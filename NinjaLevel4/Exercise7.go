package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mp := []string{"Miss", "Moneypenny", "Hellooooo James."}

	xp := [][]string{jb, mp}

	fmt.Println(xp)

	for i, v := range xp {
		fmt.Println("record: ", i)
		for j, val := range v {
			fmt.Printf("\t index position: %v \t value %v \n", j, val)
		}
	}
}
