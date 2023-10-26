package main

import (
	"fmt"
	"task1/pkg/calc"
	"task1/pkg/input"
)

func main() {
	a := input.InputNumber("Введите первое число: ")
	operator := input.ChooseOperator()
	b := input.InputNumber("Введите второе число: ")

	result := calc.Calculate(a, b, operator)
	fmt.Printf("Результат: %d %s %d = %d\n", a, operator, b, result)
}
