// Ejercicio 2 - Calcular promedio

// Un colegio necesita calcular el promedio (por alumno) de sus calificaciones.
// Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros
// y devuelva el promedio y un error en caso que uno de los números ingresados sea negativo

package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func average(numbers ...float64) (float64, error) {
	if len(numbers) == 0 {
		return 0, errors.New("no hay numeros")
	}
	var sum float64
	for _, number := range numbers {
		if number < 0 {
			return 0, errors.New("negative number")
		}
		sum += number
	}
	return sum / float64(len(numbers)), nil
}

func main() {
	numbers := make([]float64, 0)
	for _, arg := range os.Args[1:] {
		number, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Println("error at converting numbers")
		} else {
			numbers = append(numbers, number)
		}
	}
	average, err := average(numbers...)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(average)
	}
}
