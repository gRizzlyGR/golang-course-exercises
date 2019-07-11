package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	// m[`bond_james`] = []string{`Shaken, not stirred`, `Martinis`, `Women`}
	// m[`moneypenny_miss`] = []string{`James Bond`, `Literature`, `Computer Science`}
	// m[`no_dr`] = []string{`Being evil`, `Ice cream`, `Sunsets`}

	m[`goku_son`] = []string{"Fighting", "Food", "Chichi"}

	for name, favouriteThings := range m {
		fmt.Printf("%v likes:\n", name)
		for i, thing := range favouriteThings {
			fmt.Println("\t", i, thing)
		}
	}
}
