package main

import "fmt"

func main() {

	employeeSalary := map[string]float64{

		"Jack":  60000.00,
		"Blake": 120000.50,
		"Kevin": 93000.00,
	}

	fmt.Printf("Jack's Salary = %.2f\n", employeeSalary["Jack"])

	salary1, ok := employeeSalary["Kevin"]
	if ok {
		fmt.Printf("Value %.2f was found.\n", salary1)
	} else {
		fmt.Printf("The specified key was not found.\n")
	}

	if salary2, ok := employeeSalary["Jordan"]; ok {
		fmt.Printf("Value %.2f was found.\n", salary2)
		delete(employeeSalary, "Jordan")

	} else {
		fmt.Printf("The specified key was not found.\n")
	}

	_, ok = employeeSalary["Kevin"]
	if ok {
		fmt.Printf("Key exists.\n")
	} else {
		fmt.Printf("Key Doesn't exist.\n")
	}

	for name, salary := range employeeSalary {
		fmt.Println(name, salary)
	}
}
