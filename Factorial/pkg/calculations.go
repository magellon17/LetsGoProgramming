package calculations

import (
	logr "github.com/sirupsen/logrus"
)

func factorial(n int) int {
	temp := n
	result := 1

	for temp > 0 {
		result *= temp
		temp--
	}
	return result
}

func Calculate(n int, isLogged bool) (int, error) {
	if n < 0 {
		return 0, logr.Errorf("Число должно быть положительным: %d", n, logr.Fields{
			"function": "factorial",
		})
	}

	if !isLogged {
		return factorial(n), nil
	}

	logr.Info("Start calculations...")
	logr.Info("Calculate %d !", n)
	result := factorial(n)
	logr.Info("Calculations complete !")

	return result, nil
}
