package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func checkSalary(salary int) error {
	if salary < 150000 {
		return errors.New("salary is not taxable")
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
