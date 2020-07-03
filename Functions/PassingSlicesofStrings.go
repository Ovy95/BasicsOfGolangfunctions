package main

import "fmt"

func main() {
	names := []string{"Josh", "Kev", "Becks"}
	echo(names...)
}

func echo(name ...string) {
	for _, s := range name {
		fmt.Print(s, " ")
	}
}
