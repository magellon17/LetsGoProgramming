package equationUtil

import (
	"fmt"
	"os"
)

const add, sub, mult, div = "+", "-", "*", "/"

func InputNumber(prompt string) int {
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

func ChooseOperator() string {
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

func OutputResult(a, b, result int, operator string) {
	fmt.Printf("Результат: %d %s %d = %d\n", a, operator, b, result)
}
