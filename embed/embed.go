package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/eino-ext/components/embedding/ark"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func EmbedText() {
	err := godotenv.Load() // 加载环境变量
	if err != nil {
		log.Fatal("Error loading .env file") // 处理加载错误
	}

	ctx := context.Background()

	timeout := 30 * time.Second

	// 初始化嵌入器
	embedder, err := ark.NewEmbedder(ctx, &ark.EmbeddingConfig{
		APIKey:  os.Getenv("ARK_API_KEY"),
		Model:   os.Getenv("EMBEDDER"),
		Timeout: &timeout,
	})
	if err != nil {
		panic(err)
	}

	// 生成文本向量
	input := []string{
		"你好，李鑫阳",
	}

	embeddings, err := embedder.EmbedStrings(ctx, input)
	if err != nil {
		panic(err)
	}

	// 使用生成的向量
	for i, embedding := range embeddings {
		fmt.Print("文本", i+1, "的向量维度:", len(embedding))
	}
}
