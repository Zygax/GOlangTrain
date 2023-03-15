package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	for k, v := range m {
		fmt.Printf("Character: %v\t\n", k)

		for i, v2 := range v {
			fmt.Printf("\tindex: %v hobby: %v\n", i, v2)
		}
	}
}
