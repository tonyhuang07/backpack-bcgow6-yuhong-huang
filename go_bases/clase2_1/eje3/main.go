// Ejercicio 3 - Calcular salario
// Una empresa marinera necesita calcular el salario de
// sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.

// Si es categoría C, su salario es de $1.000 por hora
// Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
// Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

// Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados
// por mes y la categoría, y que devuelva su salario.

package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const (
	categoryA = "A"
	categoryB = "B"
	categoryC = "C"
)

func salaryCalculator(category string, minutes float64) (salary float64, err error) {
	if minutes < 0 {
		return 0, errors.New("minutos negativos")
	}
	horas := minutes / 60
	switch category {
	case categoryA:
		salary = horas * 3000 * 1.5
	case categoryB:
		salary = horas * 1500 * 1.2
	case categoryC:
		salary = horas * 1000
	default:
		err = errors.New("categoria invalida")
	}

	return salary, err
}

func main() {
	category := os.Args[1]
	minutes, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Println("error al convertir los minutos")
	} else {
		salary, err := salaryCalculator(category, minutes)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(salary)
		}
	}
}
