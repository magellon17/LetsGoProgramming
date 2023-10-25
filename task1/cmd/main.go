package main

import (
	"fmt"
	"os"
)

const add, sub, mult, div = "+", "-", "*", "/"

func inputNumber(prompt string) int {
	var num int
	fmt.Print(prompt)
	_, err := fmt.Scan(&num)
	for err != nil {
		fmt.Println("Некорректное число. Пожалуйста, введите числовое значение.")
		fmt.Print(prompt)
		_, err = fmt.Scan(&num)
	}
	return num
}

func chooseOperator() string {
	var operator string
	validOperators := map[string]bool{add: true, sub: true, mult: true, div: true}

	fmt.Printf("Выберите операцию (%s, %s, %s, %s): ", add, sub, mult, div)
	fmt.Fscan(os.Stdin, &operator)

	_, isValid := validOperators[operator]
	for !isValid {
		fmt.Printf("Некорректная операция. Пожалуйста, используйте символы %s, %s, %s или %s.\n", add, sub, mult, div)
		fmt.Printf("Выберите операцию (%s, %s, %s, %s): ", add, sub, mult, div)
		fmt.Fscan(os.Stdin, &operator)
		_, isValid = validOperators[operator]
	}

	return operator
}

func calculate(a, b int, operator string) int {
	switch operator {
	case add:
		return a + b
	case sub:
		return a - b
	case mult:
		return a * b
	case div:
		if b == 0 {
			panic("Ошибка: деление на ноль невозможно.")
		}
		return a / b
	default:
		panic("Некорректная операция.")
	}
}

func main() {
	a := inputNumber("Введите первое число: ")
	operator := chooseOperator()
	b := inputNumber("Введите второе число: ")

	result := calculate(a, b, operator)
	fmt.Printf("Результат: %d %s %d = %d\n", a, operator, b, result)
}
