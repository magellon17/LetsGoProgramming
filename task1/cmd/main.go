package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b int
	var operator string

	fmt.Print("Введите первое число: ")
	_, err := fmt.Scan(&a)
	for err != nil {
		fmt.Print("Некорректное число. Пожалуйста, введите числовое значение.\n")
		fmt.Print("Введите первое число: ")
		_, err = fmt.Scan(&a)
	}

	fmt.Print("Выберите операцию (+, -, *, /): ")
	fmt.Fscan(os.Stdin, &operator)
	for operator != "+" && operator != "-" && operator != "*" && operator != "/" {
		fmt.Print("Некорректная операция. Пожалуйста, используйте символы +, -, * или /.\n")
		fmt.Print("Выберите операцию (+, -, *, /): ")
		fmt.Fscan(os.Stdin, &operator)
	}

	fmt.Print("Введите второе число: ")
	_, err = fmt.Scan(&b)
	for err != nil {
		fmt.Print("Некорректное число. Пожалуйста, введите числовое значение.\n")
		fmt.Print("Введите второе число: ")
		_, err = fmt.Scan(&b)
	}

	switch operator {
	case "+":
		fmt.Printf("Результат: %d %s %d = %d", a, operator, b, a+b)
	case "-":
		fmt.Printf("Результат: %d %s %d = %d", a, operator, b, a-b)
	case "*":
		fmt.Printf("Результат: %d %s %d = %d", a, operator, b, a*b)
	case "/":
		if b == 0 {
			fmt.Printf("Ошибка: деление на ноль невозможно.")
			break
		}
		fmt.Printf("Результат: %d %s %d = %d", a, operator, b, a/b)
	}
}
