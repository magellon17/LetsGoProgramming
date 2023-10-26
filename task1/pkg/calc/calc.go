package calc

const add, sub, mult, div = "+", "-", "*", "/"

func Calculate(a, b int, operator string) int {
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
