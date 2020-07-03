package main

import "fmt"

func main() {

	employees := map[string]map[string]string{

		"BT": map[string]string{
			"firstName": "Blake",
			"lastName":  "Travis",
		},
		"JO": map[string]string{
			"firstName": "Jack",
			"lastName":  "Overton",
		},

		"WR": map[string]string{
			"firstName": "Wayne",
			"lastName":  "Rooney",
		},
	}

	if emp, ok := employees["JO"]; ok {
		fmt.Println(emp["firstName"], emp["lastName"])
	}

	for initials, emp := range employees {
		fmt.Println(initials, emp["firstName"], emp["lastName"])
	}
}
