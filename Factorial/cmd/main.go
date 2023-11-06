package main

import (
	"flag"
	"os"
	"strconv"

	calc "LetsGoProgramming/Factorial/pkg/calculations"

	logr "github.com/sirupsen/logrus"
)

func main() {
	loggingFlag := flag.Bool("log", false, "Enable logging")
	flag.Parse()

	args := os.Args[1:]
	if len(args) != 1 {
		logr.Error("Необходимо указать один аргумент: число для факториала")
		return
	}

	n, err := strconv.Atoi(args[0])
	if err != nil {
		logr.Errorf("Не удалось преобразовать аргумент в число: %v", err)
		return
	}

	result, err := calc.Calculate(n, loggingFlag)
	if err != nil {
		logr.Error(err)
		return
	}

	logr.Info(result)
}
