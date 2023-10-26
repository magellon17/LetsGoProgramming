package main

import (
	"task1/pkg/calc"
	"task1/pkg/io"
)

func main() {
	a := io.InputNumber("Введите первое число: ")
	operator := io.ChooseOperator()
	b := io.InputNumber("Введите второе число: ")

	result := calc.Calculate(a, b, operator)
	io.OutputResult(a, b, result, operator)
}
