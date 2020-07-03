package main

import "fmt"

func main(){
//value
x:= 1
//Pointers
p:= &x
q:= &x

		fmt.Printf("x => type =%T value=%d \n",x,x)
		fmt.Println("Just the standard value of X so its 1")
		println(" ")
		fmt.Printf("&x => type =%T value=%x \n",&x,&x)
		println("&x This is the address value of x ")
		println(" - - - - - - - - - - - - - - - - - - - - ")
		fmt.Printf("*p => type =%T value=%d \n",*p,*p)
println("...*p This is the pointer to the value so this takes the value of x so this is 1")
println("")
		fmt.Printf("p => type =%T value=%x \n",p,p)
println("... This is equal to the value's ADDRESS of x ")
println("")
		fmt.Printf("&p => type =%T value=%x \n",&p,&p)
		println("... This is equal to the value's ADDRESS of p its own address ")
		println(" - - - - - - - - - - - - - - - - - - - - ")
		println(" ")
		fmt.Printf("*q => type =%T value=%d \n",*q,*q)
		println("...*q This is the pointer to the value so this takes the value of x so this is 1")
println("")
		fmt.Printf("q => type =%T value=%x \n",q,q)
		println("... This is equal to the value's ADDRESS of x ")
		println("")
		fmt.Printf("&q => type =%T value=%x \n",&q,&q)
		println("This is equal to the value's ADDRESS of q its own address ")
		println("")
}