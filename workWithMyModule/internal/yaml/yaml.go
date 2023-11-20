package yaml

import (
	"os"

	//log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type File struct {
	Filename  string `yaml:"filename"`
	Substring string `yaml:"substring"`
}

func Unmarshal(yamlPath string) ([]File, error) {
	//log.Info("Start reading...")

	yamlFile, err := os.ReadFile(yamlPath)
	if err != nil {
		return nil, err
	}

	var files struct {
		Files []File
	}

	err = yaml.Unmarshal(yamlFile, &files)
	if err != nil {
		return nil, err
	}

	return files.Files, nil
}
