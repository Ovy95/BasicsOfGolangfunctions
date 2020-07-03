package main

import "fmt"

func main() {
	const (
		off, false, no = iota, iota, iota
		yes, on, true
	)

	fmt.Printf("off=%d false=%d no=%d yes=%d true=%d on=%d\n", off, false, no, yes, true, on)
}
