package main

import (
	"fmt"
)

// Exercise : Calculate average of three integerse

func main() {

	a := 5
	b := 10
	c := 15

	fmt.Println((a + b + c) / 3)
	fmt.Printf("Average of %d, %d, and %d is %d.\n",
		a, b, c, (a+b+c)/3)
}
