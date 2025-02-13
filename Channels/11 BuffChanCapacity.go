package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string, 3)
	fmt.Printf("capacity:%d length:%d\n", cap(c), len(c))

	c <- "Message 1"
	c <- "Message 2"

	fmt.Printf("capacity:%d length:%d\n", cap(c), len(c))

	c <- "Message 3"
	fmt.Printf("capacity:%d length:%d\n", cap(c), len(c))

	time.Sleep(5 * time.Second)
	fmt.Println(<-c)
	time.Sleep(2 * time.Second)
	fmt.Println(<-c)

	fmt.Printf("capacity:%d length:%d\n", cap(c), len(c))
}
