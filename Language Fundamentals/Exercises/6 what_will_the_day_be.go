package main

import "fmt"

func main() {
	// Sunday =0 Monday= 1 Tuesday = 2 Wednesday = 3
	// Thursday = 4 Friday = 5 Saturday = 6 Sunday = 7
	today := 6
	duration := 8
	futureDay := (today + duration) % 7

	fmt.Printf("today=%d, duration=%d future day=%d\n",
		today, duration, futureDay)

}
