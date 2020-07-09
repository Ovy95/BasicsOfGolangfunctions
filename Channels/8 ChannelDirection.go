/* ASSIGNEMNT -Assuming your current bank balance is 200, write a program containng these functions:
1) credit () function with a bi-directional channel parameter that sends random 
numbers (between 1 and 10) to the channel that will be added to the balance 
2) debit() function with a send-only channel parameter that sends random numbers between -10 and -1
to the channel that will be subtracted from thr balance 
3) balance() function with a 'receive-only channel' parameter that adds/subtracts the random values (generated in steps 1 & 2)to the balance

Here a sample output that indicates that the balance is ramdomly changing
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var currentBalance = 200

func main() {
	fmt.Printf("\nBANK IS NOW OPEN\n Press the enter Key to stop the program")
	rand.Seed(time.Now().UnixNano()) // Used for the random numbers can be used within functions 

	c := make(chan int)
	go credit(c)
	go debit(c)
	go balance(c)
	
	var input string
	fmt.Scanln(&input)
}

func credit( c chan int) { //Bidrectional channel  could also be a send only
	fmt.Println("(f) entering Creditmode()")

	for i := 0; ; i++ {
	c <- rand.Intn(9) + 1
	}
}

func debit( c chan<- int) { //Send only channel 
	for i := 0; ; i++ {
	c <- rand.Intn(9) - 10
	}
}

func balance( c <- chan int) { // Receive only channel
	for{
		num,ok := <-c
		if ok == false {
			fmt.Println("Error :")
			break
		}

		oldBalance := currentBalance
		currentBalance += num

		fmt.Printf("=> %d + (%d) = %d \n",oldBalance , num , currentBalance)
		time.Sleep(1 * time.Second)
	}

}