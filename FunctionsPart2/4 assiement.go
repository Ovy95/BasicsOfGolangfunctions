 package main
import "fmt"
func main() {
	addCounter,	multCounter := addby(),multBy()
	
	fmt.Println(addCounter(10)," ")
	fmt.Println(addCounter(10)," ")
	fmt.Println(addCounter(10)," ")
	fmt.Println(multCounter(2)," ")
	fmt.Println(multCounter(10)," ")
}

func addby()  func(int) int {
	total := 0
	return func(i int) (ret int) {
	total += i
	ret = total
	return
	}
}

func multBy()  func(int) int {
	total := 1
	return func(i int) (ret int) {
	total *= i
	ret = total
	return
	}
}