package main

import (
	"fmt"
	"os"
	"strconv"
)

const MinumumHours = 80

func calculateSalary(hourlyRate int, monthlyHours ...int) (salaries []int, err error) {
	var insuficientHours []int
	for _, hours := range monthlyHours {
		if hours < MinumumHours {
			insuficientHours = append(insuficientHours, hours)
		} else {
			salary := hourlyRate * hours
			if salary > 150000 {
				salary = salary - (salary * 10 / 100)
			}
			salaries = append(salaries, salary)
		}

	}
	if len(insuficientHours) > 0 {
		err = fmt.Errorf("insuficient hours are : %v", insuficientHours)
	}

	return salaries, err
}
func caulateBonus(salaries []int) (bonus int, err error) {
	var negativeSalaries []int

	var maxSalary int
	if len(salaries) == 0 {
		err = fmt.Errorf("no salary to calculate bonus")
		return 0, err
	}
	maxSalary = salaries[0]
	for _, salary := range salaries {
		if salary < 0 {
			negativeSalaries = append(negativeSalaries, salary)
		} else {
			if salary > maxSalary {
				maxSalary = salary
			}
		}
	}
	if len(negativeSalaries) > 0 {
		err = fmt.Errorf("negative salaries are : %v", negativeSalaries)
	}
	bonus = maxSalary / 12 * len(salaries)

	return bonus, err
}

func main() {
	hourlyRate, _ := strconv.Atoi(os.Args[1])
	monthlyHours := make([]int, 0)
	for _, arg := range os.Args[2:] {
		monthlyHour, _ := strconv.Atoi(arg)
		monthlyHours = append(monthlyHours, monthlyHour)
	}
	salaries, err := calculateSalary(hourlyRate, monthlyHours...)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("calculated salaries are:", salaries)

	bonus, err := caulateBonus(salaries)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("calculated bonus is:", bonus)
}
