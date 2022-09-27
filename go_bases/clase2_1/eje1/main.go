// Ejercicio 1 - Impuestos de salario
// Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo,
// para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario.
// Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17% del sueldo y
// si gana más de $150.000 se le descontará además un 10%.

package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func taxCalculator(salary float64) (float64, error) {

	if salary < 0 {
		return 0, errors.New("negative salary")
	}

	if salary < 150000 && salary >= 50000 {
		return salary * 0.17, nil
	}

	if salary >= 150000 {
		return salary * (0.17 + 0.1), nil
	}
	return 0, nil
}

func main() {
	salary, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Error at converting numbers")
	} else {
		tax, err := taxCalculator(salary)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(tax)
		}
	}
}
