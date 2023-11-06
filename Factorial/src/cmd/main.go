package main

import (
	"flag"
	"go/LetsGoProgramming/Factorial/src/pkg/calculations"
	"log"
	"os"
	"strconv"

	logr "github.com/sirupsen/logrus"
)

func main() {
	loggingFlag := flag.Bool("log", false, "Enable logging")
	flag.Parse()

	args := os.Args[1:]
	if len(args) < 1 {
		logr.Fatal("Необходимо указать число для факториала")
		return
	}

	argvNumI := 0
	if len(args) > 1 {
		argvNumI = 1
	}

	n, err := strconv.Atoi(args[argvNumI])
	if err != nil {
		logr.Fatalf("Не удалось преобразовать аргумент в число: %v", err)
		return
	}

	result, err := calculations.Calculate(n, *loggingFlag)
	if err != nil {
		logr.Fatal(err)
		return
	}

	log.Printf("Результат вычисления факториала числа %v это %v", n, result)
}
