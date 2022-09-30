package main

import (
	"fmt"
	"os"
	"strconv"
)

type NotTaxableError struct {
	Msg string
}

func (e *NotTaxableError) Error() string {
	return e.Msg
}

func checkSalary(salary int) error {
	if salary < 150000 {
		return &NotTaxableError{
			Msg: "salary is not taxable",
		}
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
