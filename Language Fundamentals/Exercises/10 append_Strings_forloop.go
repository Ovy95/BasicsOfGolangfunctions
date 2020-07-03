package main

import "fmt"

func main() {

	Ja := "JA"
	Ck := "Ck"

	var myName []rune

	for _, n := range Ja {
		myName = append(myName, n)
	}

	for _, n := range Ck {
		myName = append(myName, n)
	}

	fmt.Printf("%q\n", myName)
	fmt.Printf("%q\n", string(myName))
}
