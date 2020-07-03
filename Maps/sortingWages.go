package main

import (
	"fmt"
	"sort"
)

func main() {

	wages := map[string]float64{

		"Jack":  60000.00,
		"Blake": 120000.50,
		"Kevin": 93000.00,
	}

	names := make([]string, 0, len(wages))
	for n := range wages {
		names = append(names, n)
	}
	fmt.Println(names)

	sort.Strings(names)
	fmt.Println(names)

	for _, n := range names {
		fmt.Printf("%s\t%.2f\n", n, wages[n])
	}

}
