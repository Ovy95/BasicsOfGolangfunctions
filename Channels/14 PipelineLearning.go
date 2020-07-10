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

func main()  {
	babbler := tjarratt.NewBabbler()
	newWords := make(chan string)
	uWords := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			newWords <- babbler.Babble()
		}
		close(newWords)
	}()

	go func()  {
		for w := range newWords {
			uWords <- w + "- - > " + strings.ToUpper(w)
		}
		close(uWords)
	}()

	for w := range uWords {
		fmt.Println(w)
	}
	
}


