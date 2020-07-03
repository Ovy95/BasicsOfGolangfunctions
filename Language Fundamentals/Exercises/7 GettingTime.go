package main

import (
	"fmt"
	"time"
)

func main() {
	toatlMilliSeconds := time.Now().UnixNano() / int64(time.Millisecond)

	totalSeconds := toatlMilliSeconds / 1000
	currentSecond := totalSeconds % 60
	totalMinutes := totalSeconds / 60
	currentMintue := totalMinutes % 60
	totalHours := totalMinutes / 60
	currentHour := totalHours%24 + 1

	fmt.Printf("toatlMilliSeconds = %d \n", toatlMilliSeconds)
	fmt.Printf("currentSecond = %d \n", currentSecond)
	fmt.Printf("totalMinutes = %d \n", totalMinutes)
	fmt.Printf("currentMintue = %d \n", currentMintue)
	fmt.Printf("totalHours = %d \n", totalHours)
	fmt.Printf("currentHour = %d \n", currentHour)

	fmt.Printf("The current time is %d:%d ", currentHour, currentMintue)
}
