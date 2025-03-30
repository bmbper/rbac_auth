package main

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

var RbacConfig struct {
	AppName    string `yaml:"app_name"`
	ServerPort string `yaml:"server_port"`
}

func LoadConfig() error {
	file, err := os.Open("api/config.yaml")
	if err != nil {
		return fmt.Errorf("failed to open YAML file: %v", err)
	}
	defer file.Close()
	yamlFile, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read YAML file: %v", err)
	}
	err = yaml.Unmarshal(yamlFile, &RbacConfig)
	if err != nil {
		return fmt.Errorf("failed to unmarshal YAML data: %v", err)
	}
	return nil
}
