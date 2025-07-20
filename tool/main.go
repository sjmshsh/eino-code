package main

import (
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load() // 加载环境变量
	if err != nil {
		log.Fatal("Error loading .env file") // 处理加载错误
	}

	SimpleAgent()
}
