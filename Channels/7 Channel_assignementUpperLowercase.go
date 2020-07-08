// ASSIGNEMNT - Write a program that contains two channels that
// Accept a text and convert the text to uppercase and lowercase.

package main

func main() {
	text := "Please Modify This Text!"

	cUpper := make(chan string)
	clower := make(chan string)

	go uppercase(text, cUpper)
}

func uppercase(s string, c chan string) {
	c <- string.ToUpper(s)
}

func lowercase(s string, c chan string) {
	c <- string.ToLower(s)
}
