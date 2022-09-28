package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Selection Sort
func selectionsort(array []int) {
	fmt.Println("start selection sort with array of size: ", len(array))
	var n = len(array)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if array[j] < array[minIdx] {
				minIdx = j
			}
		}
		array[i], array[minIdx] = array[minIdx], array[i]
	}

}

// Bubble Sort
func BubbleSort(array []int) {
	fmt.Println("start bubble sort with array of size: ", len(array))
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}

// Insertion Sort
func insertionsort(array []int) {
	fmt.Println("start insertion sort with array of size: ", len(array))
	var n = len(array)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if array[j-1] > array[j] {
				array[j-1], array[j] = array[j], array[j-1]
			}
			j = j - 1
		}
	}
}

func selectSort(array1 []int, array2 []int, array3 []int, functionSort func([]int)) time.Duration {
	start := time.Now()
	functionSort(array1)
	functionSort(array2)
	functionSort(array3)
	return time.Since(start)
}

func main() {
	array1 := rand.Perm(100)
	array2 := rand.Perm(1000)
	array3 := rand.Perm(10000)

	timeSelectSort := make(chan time.Duration)
	timeBubbleSort := make(chan time.Duration)
	timeInsertionSort := make(chan time.Duration)

	go func() {
		timeSelectSort <- selectSort(array1, array2, array3, selectionsort)
	}()

	go func() {
		timeBubbleSort <- selectSort(array1, array2, array3, BubbleSort)
	}()

	go func() {
		timeInsertionSort <- selectSort(array1, array2, array3, insertionsort)
	}()

	fmt.Println("time of selection sort: ", <-timeSelectSort)
	fmt.Println("time of bubble sort: ", <-timeBubbleSort)
	fmt.Println("time of insertion sort: ", <-timeInsertionSort)

}
