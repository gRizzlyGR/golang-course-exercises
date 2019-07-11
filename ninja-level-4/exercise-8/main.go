package main

import (
	"fmt"
)

func main() {
	jb := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	// jb[`bond_james`] = []string{`Shaken, not stirred`, `Martinis`, `Women`}
	// jb[`moneypenny_miss`] = []string{`James Bond`, `Literature`, `Computer Science`}
	// jb[`no_dr`] = []string{`Being evil`, `Ice cream`, `Sunsets`}

	for name, favouriteThings := range jb {
		fmt.Printf("%v likes:\n", name)
		for i, thing := range favouriteThings {
			fmt.Println("\t", i, thing)
		}
	}
}
