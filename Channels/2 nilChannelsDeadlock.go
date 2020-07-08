package main

import "fmt"

func main() {
	// Two ways of declaring channels
	shortWay := make(chan string)
	// Delaring & defining using short hand declaration

	var c chan int //Long way
	//If statement is created to stop nil channels
	if c == nil {
		c = make(chan int)
	}

	fmt.Printf("Type of c: %T", c)
	fmt.Println()
	fmt.Printf("Type of Shortway: %T", shortWay)
	close(c)
	close(shortWay)

	deadlockChannel := make(chan string)
	deadlockChannel <- "No one likes my channel!" // fatel error : all go routines are asleep - deadlock !
	fmt.Println("This demonstrates a go routine that is sending data on a channel, but no other routines will receive the data. Creating the error above")
}
