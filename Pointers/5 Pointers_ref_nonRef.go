package main

import "fmt"

func main() {
	x := 10

	fmt.Printf("x=%d\n", x)
	f1(x)
	fmt.Printf("x=%d\n", x)
	fmt.Println("......")
	fmt.Printf("x=%d\n", x)
	f2(&x)
	fmt.Printf("x=%d\n", x)

}

func f1(y int) {
	fmt.Printf("(f1) y=%d\n", y)
	y += 2
	fmt.Printf("(f1) y=%d\n", y)
}

func f2(y *int) {
	fmt.Printf("(f2) y=%d\n", y)
	*y += 2
	fmt.Printf("(f2) y=%d\n", y)
}
