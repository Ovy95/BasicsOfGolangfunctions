package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 1)

	for i := 1; i <= 10; i++ {
		go printMsg(c, i)
	}
	time.Sleep(10 * time.Second)
}

func printMsg(c chan int, id int) {
	fmt.Printf("ooo %d is waiting for a channel space\n", id)

	c <- id
	fmt.Printf("=== %d has a channel space\n", id)
	time.Sleep(6000 * time.Microsecond)
	fmt.Printf("xxx %d has released\n", id)
	<-c
}
