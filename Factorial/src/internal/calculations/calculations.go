package calculations

import (
	"errors"

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
		return 0, errors.New("factorial error: negative number")
	}

	if !isLogged {
		return factorial(n), nil
	}

	logr.Info("Start calculations...")
	logr.Infof("Calculate %d !", n)
	result := factorial(n)
	logr.Info("Calculations complete !")

	return result, nil
}
