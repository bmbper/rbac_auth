package base

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type AppConfig struct {
	Name string `yaml:"app_name"`
	Code string `yaml:"app_code"`
	Icon string `yaml:"app_icon"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db_name"`
	PoolSize int    `yaml:"pool_size"`
}

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	App      AppConfig      `yaml:"app"`
	Database DatabaseConfig `yaml:"database"`
}

var GlobalConfig Config

func LoadAppConfig(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("读取配置文件失败: %v", err)
	}
	defer file.Close()

	yamlFile, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("读取配置文件失败: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &GlobalConfig)
	if err != nil {
		return fmt.Errorf("提取配置信息失败: %v", err)
	}

	return nil
}
