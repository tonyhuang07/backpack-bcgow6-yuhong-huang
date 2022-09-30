package main

import (
	"fmt"
	"os"
	"strconv"
)

func checkSalary(salary int) error {
	if salary < 150000 {
		return fmt.Errorf("salary is not taxable, salary is %v", salary)
	}
	return nil
}

func main() {
	salary, _ := strconv.Atoi(os.Args[1])
	error := checkSalary(salary)
	if error != nil {
		fmt.Println("error:", error)
	} else {
		fmt.Println("need to pay taxes")
	}
}
