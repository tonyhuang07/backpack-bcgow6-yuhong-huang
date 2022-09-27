// Ejercicio 4 - Calcular estadísticas

// Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas
// de calificaciones de los alumnos de un curso, requiriendo calcular los valores mínimo,
// máximo y promedio de sus calificaciones.

// Se solicita generar una función que indique qué tipo de cálculo se quiere realizar
// (mínimo, máximo o promedio) y que devuelva otra función
// ( y un mensaje en caso que el cálculo no esté definido) que se le puede pasar una
// cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior

package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func calculateMin(numbers ...float64) float64 {
	min := numbers[0]
	for _, number := range numbers {
		if number < min {
			min = number
		}
	}
	return min
}

func calculateMax(numbers ...float64) float64 {
	max := numbers[0]
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

func calculateAvg(numbers ...float64) float64 {
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

func operation(operation string) (function func(...float64) float64, err error) {
	switch operation {
	case minimum:
		function = calculateMin
	case average:
		function = calculateAvg
	case maximum:
		function = calculateMax
	default:
		err = errors.New("invalid operation")
	}

	return function, err
}

func main() {
	var opcode string = os.Args[1]
	operation, err := operation(opcode)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("operation:", opcode)
		numbers := make([]float64, 0)
		for _, arg := range os.Args[2:] {
			number, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Println("error at converting numbers")
			} else {
				numbers = append(numbers, number)
			}
		}
		fmt.Println(operation(numbers...))
	}
}
