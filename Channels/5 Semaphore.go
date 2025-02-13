package main

import "fmt"

var wordSet1 = []string{"small", "medium", "large"}
var wordSet2 = []string{"red", "green", "blue", "Yellow"}

var c = make(chan string)
var b = make(chan bool)

func main() {

	f := []interface{}{sender1, sender2}

	for i := range f {
		go f[i].(func())() // (func(type))(value)
	}
	go closeSenders()

	fmt.Println("Before getting to the 'channel for range loop'.")

	for val := range c {
		fmt.Println(val)
	}
}
func sender1() {
	for _, w := range wordSet1 {
		c <- w
	}
	b <- true
}

func sender2() {
	for _, w := range wordSet2 {
		c <- w
	}
	b <- true
}

func closeSenders() {
	<-b
	<-b
	close(c)
}
