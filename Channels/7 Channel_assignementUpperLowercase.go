// ASSIGNEMNT - Write a program that contains two channels that
// Accept a text and convert the text to uppercase and lowercase.

package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	text := "Please Modify This Text!"
	cUpper := make(chan string)
	clower := make(chan string)

	fmt.Printf("\ntext before uppercase \n  %s\n \n", text)
	go uppercase(text, cUpper)
	go lowercase(text, clower)
	sUpper, sLower := <-cUpper, <-clower
	fmt.Printf("sUpper=%s \n", sUpper)

	fmt.Printf("slower=%s \n", sLower)

}

func uppercase(s string, c chan string) {
	fmt.Println("(f) entering uppercase()")
	time.Sleep(3 * time.Second)
	c <- strings.ToUpper(s)
}

func lowercase(s string, c chan string) {
	fmt.Println("(f) entering lowercase()")
	time.Sleep(3 * time.Second)
	c <- strings.ToLower(s)
}
