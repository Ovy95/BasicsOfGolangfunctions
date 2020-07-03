package main

import "fmt"

// First type of for loop single Condition
func main() {
	IncrementForLoop()
	rangeclauseLoop()
	SingleCondition()
	nestedLoop()
}
func SingleCondition() {
	//condition
	a := 1
	b := 2
	for a < b {
		//Body
		a *= 2
		fmt.Println("single for loop have been run")
	}
}

// With Init and post statement
// 1) initialize variable "i" with zero;
// 2) check the condition; execute the body lines; increase "i" by 1;
// 3) repeat step 2 until i < 5
func IncrementForLoop() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		fmt.Println("Increment for loop have been run")
	}
}

func rangeclauseLoop() {
	x := [...]int{10, 20, 30} // array of integers
	for i, val := range x {   //  i v   i v   i  v
		fmt.Println(i, val) // 0 10, 1 20, 2 30
		fmt.Println("range Clause for loop have been run")
		//   v  i   v i   v i
		fmt.Println(val, i) // 10 0   20 1  30 2
	}
}

func nestedLoop() {
	for i := 1; i < 5; i++ {
		fmt.Printf("\ni=%d ** ", i)
		for j := 1; j < 3; j++ {
			fmt.Printf("%d ", i*j)
		}
	}

}
