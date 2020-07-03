package main

import "fmt"

func main() {

	Fahrenheit := 100.00
	Celisius := ((Fahrenheit - 32.0) * 5.0 / 9.0)
	fmt.Println(Celisius)

	fmt.Printf("Fahrenheit=%f ==> Celisius=%f\n",
		Fahrenheit, Celisius)
}
