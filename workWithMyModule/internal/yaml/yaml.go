package yaml

import (
	"os"

	"gopkg.in/yaml.v3"
)

type File struct {
	Path      string `yaml:"filename"`
	Substring string `yaml:"substring"`
}

func Unmarshal(yamlPath string) ([]File, error) {
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
