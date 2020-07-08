// The problem with this is that one of the channels is going to close first not allowing both Wordsets
// to be executed which will result in a Panic error like shown if tried running the program
// See 5 Semaphore.go on how this issue is solved

package main

import "fmt"

var wordSet1 = []string{"small", "medium", "large"}
var wordSet2 = []string{"red", "green", "blue", "Yellow"}

func main() {
	c := make(chan string)

	go queue1(c)
	go queue2(c)

	// time.Sleep(2 * time.Second)
	for val := range c {
		fmt.Println(val)
	}
}

func queue1(c chan string) {
	for _, w := range wordSet1 {
		c <- w
	}
	close(c) //unpredicted behaviour !
}
func queue2(c chan string) {
	for _, w := range wordSet2 {
		c <- w
	}
	close(c) //unpredicted behaviour !
}
