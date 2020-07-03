package main

import "fmt"

func main() {

	power2 := make(map[int]int)

	power2[2] = 4
	power2[3] = 9

	fmt.Println(power2)
	fmt.Println(power2[2])
	fmt.Println(power2[3])
	fmt.Println(power2[1])

}
