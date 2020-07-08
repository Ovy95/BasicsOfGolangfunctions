package main

import (
	"fmt"
	"sync"
)

var waitG sync.WaitGroup

func main() {
	c := make(chan int)

	waitG.Add(2)
	go send(c)
	go receive(c)
	waitG.Wait()

}

func send(c chan int) {
	for i := 1; i < 6; i++ {
		c <- i
	}
	waitG.Done()
}

func receive(c chan int) {
	for i := 1; i < 6; i++ {

		fmt.Println(<-c)
	}
	waitG.Done()
}
