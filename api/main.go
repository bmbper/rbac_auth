package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

func main() {
	// 读取配置文件
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	var config struct {
		AppName    string `yaml:"app_name"`
		ServerPort string `yaml:"server_port"`
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}
	// 初始化Gin路由
	r := gin.Default()
	r.Run(":" + config.ServerPort)
}
