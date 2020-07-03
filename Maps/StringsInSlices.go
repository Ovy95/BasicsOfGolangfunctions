/*
Create a slice of string , including some repating words
Write a program to remove the repeating words
*/

package main

import "fmt"

func main() {
	s :=[]string{"one","two","two","three","four","four","one","four"}

	m := make(map[string]bool)

	for i := range s {
		word := s[i]
		if !m[word] {
			m[word]= true
		} 
	}

	fmt.Println(s)
	fmt.Println(m)


}