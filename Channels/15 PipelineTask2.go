/*
Write a program including the following go routines:
1 ) 'newWords' go routine that produces random workds and send them to a channel;
2) The 'uwords' (Uppercase words ) goroutine that converts the randomly-generated words to uppercase
Use anonymous functions to write your goroutines
*/
package main

import (
	"fmt"
	"strings"

	tjarratt "github.com/tjarratt/babble"
)

func main() {
	//
	newWords := make(chan string)
	uWords := make(chan string)

	go sendWords(newWords)
	go convertWords(uWords, newWords)
	printWords(uWords)
}

func sendWords(out chan<- string) {
	babbler := tjarratt.NewBabbler()
	for i := 0; i < 5; i++ {
		out <- babbler.Babble()
	}
	close(out)
}

func convertWords(out chan<- string, in <-chan string) {
	for w := range in {
		out <- w + " --> " + strings.ToUpper(w)
	}
	close(out)
}

func printWords(in <-chan string) {
	for w := range in {
		fmt.Println(w)
	}
}
