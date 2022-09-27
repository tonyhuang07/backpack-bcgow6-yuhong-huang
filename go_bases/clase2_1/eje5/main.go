/*
Ejercicio 5 - Calcular cantidad de alimento

Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas.
Por el momento solo tienen tarántulas, hamsters, perros, y gatos, pero se espera que puedan
haber muchos más animales que refugiar.

perro necesitan 10 kg de alimento
gato 5 kg
Hamster 250 gramos.
Tarántula 150 gramos.

Se solicita:
Implementar una función Animal que reciba como parámetro un valor de tipo texto con el
animal especificado y que retorne una función y un mensaje (en caso que no exista el animal)
Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del
tipo de animal especificado.

*/

package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const (
	Dog     = "dog"
	Cat     = "cat"
	Hamster = "hamster"
	Spider  = "spider"
)

func dogFoodCalculator(dogs float64) (float64, error) {
	if dogs < 0 {
		return 0, errors.New("negative number")
	}
	return dogs * 10, nil
}

func catFoodCalculator(cats float64) (float64, error) {
	if cats < 0 {
		return 0, errors.New("negative number")
	}
	return cats * 5, nil
}

func hamsterFoodCalculator(hamsters float64) (float64, error) {
	if hamsters < 0 {
		return 0, errors.New("negative number")
	}
	return hamsters * 0.25, nil
}

func spiderFoodCalculator(spiders float64) (float64, error) {
	if spiders < 0 {
		return 0, errors.New("negative number")
	}
	return spiders * 0.15, nil
}

func Animal(animal string) (function func(float64) (float64, error), err error) {
	switch animal {
	case Dog:
		function = dogFoodCalculator
	case Cat:
		function = catFoodCalculator
	case Hamster:
		function = hamsterFoodCalculator
	case Spider:
		function = spiderFoodCalculator
	default:
		err = errors.New("type of animal not found")
	}
	return function, err
}

func main() {
	var typeAnimal string = os.Args[1]
	quantity, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Println("error at converting numbers")
	}

	function, err := Animal(typeAnimal)
	if err != nil {
		fmt.Println(err)
	} else {
		food, err := function(quantity)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("cant of food:", food)
		}
	}
}
