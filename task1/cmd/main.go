package main

import (
	"task1/pkg/calc"
	util "task1/pkg/equationUtil"
)

func main() {
	a := util.InputNumber("Введите первое число: ")
	operator := util.ChooseOperator()
	b := util.InputNumber("Введите второе число: ")

	result := calc.Calculate(a, b, operator)
	util.OutputResult(a, b, result, operator)
}
