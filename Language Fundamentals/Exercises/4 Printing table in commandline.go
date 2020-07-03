package main

import "fmt"

func main() {
	i := 2

	fmt.Println("i \t i*2 \t i*3")
	fmt.Println("== \t === \t ===")

	fmt.Printf("%d \t %d \t %d\n", i, i*2, i*3)
	i++
	fmt.Printf("%d \t %d \t %d\n", i, i*2, i*3)
	i++
	fmt.Printf("%d \t %d \t %d\n", i, i*2, i*3)
}
