package main

import "fmt"

func main() {
	//value
	x := []int{2, 4, 6, 8, 10}
	//Pointers
	p := &x

	fmt.Printf("x=%v *p=d &x???=%x p???=%x %p=%x\n",
		x, *p, &x, p, &p) // &x and p are meaningless

	p0 := &x[0]
	fmt.Printf("x[0]=%v *p0=d &x[0]=%x p0=%x %p0=%x\n",
		x[0], *p, &x[0], p, &p0)

	p2 := &x[2]
	fmt.Printf("x[2]=%v *p2=d &x[2]=%x p2=%x %p2=%x\n",
		x[2], *p, &x[2], p, &p2)

	fmt.Println()
	y := x[2:] // back end of the slices soo [6,8,10]
	q := &y
	fmt.Printf("*q=%d &q=%x\n", *q, &q)

	q0 := &y[0]
	fmt.Printf("y[0]=%v *q0=%d &y[0]=%x q0=%x %q0=%x\n",
		y[0], *q0, &y[0], q0, &q0)

	(*p)[4] = 20
	//DE refer this changes the last value to 20 in the slices of Intergers
	fmt.Println()
	for i := range x {
		fmt.Println(i, x[i], (*p)[i])
	}
	fmt.Println()
	for i := range y {
		fmt.Println(i, y[i], (*q)[i])
	}
}
