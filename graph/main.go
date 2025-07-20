package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load() // 加载环境变量
	if err != nil {
		log.Fatal("Error loading .env file") // 处理加载错误
	}
	ctx := context.Background()
	input := map[string]string{
		"role":    "tsundere",
		"content": "你好啊",
	}
	answer := OrcGraphWithModel(ctx, input)

	fmt.Println(answer.Content)
}
