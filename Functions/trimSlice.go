package main

import "fmt"

func main() {

	data := []string{"Daisy", "Rose", "", "Steve"}

	fmt.Printf("%q\n", trimSlice(data))
	fmt.Print("%q\n", data)
}

func trimSlice(data []string) []string {

	var newData []string
	for _, d := range data {
		if d != "" {
			newData = append(newData, d)
		}

	}
	return newData
}
