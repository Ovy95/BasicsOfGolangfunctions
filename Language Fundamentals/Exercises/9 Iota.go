package main

import "fmt"

func main() {

	const (
		_   = iota
		Ten = 10 * iota
		Twenty
		Thirty
	)

	fmt.Printf("Ten=%d Twenty=%d Thirty%d\n", Ten, Twenty, Thirty)

}
