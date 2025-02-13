package main

import "fmt"

func main() {

	nextPos1 := getPositiveInt()

	fmt.Println(nextPos1())
	fmt.Println(nextPos1())
	fmt.Println(nextPos1())

	nextPos2 := getPositiveInt()
	fmt.Println(nextPos2())

	fmt.Printf("'%T' '%T' \n ", nextPos1, nextPos1())
	fmt.Printf("'%x' '%x' \n ", &nextPos1, &nextPos2)
}

func getPositiveInt() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}
