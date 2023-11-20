package main

import (
	"LetsGoProgramming/workWithMyModule/internal/yaml"
	"fmt"
	"os"

	mod "github.com/magellon17/testmod/contains"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.WarnLevel)

	if len(os.Args) < 2 {
		log.Fatal("The args must contain yml file path")
	}

	ymlPath := os.Args[1]
	files, err := yaml.Unmarshal(ymlPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		contains, err := mod.Contains(file.Filename, file.Substring)
		if err != nil {
			log.Warn(err)
			continue
		}

		fmt.Printf("File '%s' contains %s : %v\n", file.Filename, file.Substring, contains)
	}
}
