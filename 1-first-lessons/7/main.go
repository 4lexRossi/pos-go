package main

import "fmt"

func main() {
	salaryMap := map[string]int{"alex": 1000, "jon": 2000, "xpto": 3000}
	delete(salaryMap, "jon")
	salaryMap["jon"] = 500

	// sal := make(map[string]int)
	// sal1 := map[string]int{}

	for name, salary := range salaryMap {
		fmt.Printf("the salary of %s is %d\n", name, salary)
	}

	for _, salary := range salaryMap {
		fmt.Printf("the salary is %d\n", salary)
	}
}
