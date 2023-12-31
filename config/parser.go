package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func ReadConfig(config *AppConfig, filename string) {
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	if err = yaml.Unmarshal(yamlFile, &config); err != nil {
		log.Fatal(err)

	}
}
