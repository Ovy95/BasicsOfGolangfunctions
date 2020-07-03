package main

import "fmt"

func main() {

	const PI = 3.14159
	var radius float64 = 3

	area := radius * radius * PI

	fmt.Printf("The area of the circle of radius %5.2f is %5.2f.\n",
		radius, area)
}
