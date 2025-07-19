package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/eino-ext/components/embedding/ark"
	"github.com/cloudwego/eino-ext/components/retriever/milvus"
	"github.com/joho/godotenv"
	cli "github.com/milvus-io/milvus-sdk-go/v2/client"
	"log"
	"os"
	"time"
)

var MilvusCli cli.Client

func Init() {
	//初始化客户端
	ctx := context.Background()
	client, err := cli.NewClient(ctx, cli.Config{
		Address: "localhost:19530",
		DBName:  "Eino",
	})
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	MilvusCli = client
}

func main() {
	Init()
	err := godotenv.Load() // 加载环境变量
	if err != nil {
		log.Fatal("Error loading .env file") // 处理加载错误
	}

	ctx := context.Background()
	// 初始化嵌入器
	timeout := 30 * time.Second
	embedder, err := ark.NewEmbedder(ctx, &ark.EmbeddingConfig{
		APIKey:  os.Getenv("ARK_API_KEY"),
		Model:   os.Getenv("EMBEDDER"),
		Timeout: &timeout,
	})
	if err != nil {
		panic(err)
	}

	retriever, err := milvus.NewRetriever(ctx, &milvus.RetrieverConfig{
		Client:      MilvusCli,
		Collection:  "test",
		Partition:   nil,
		VectorField: "vector",
		OutputFields: []string{
			"id",
			"content",
			"metadata",
		},
		TopK:      1,
		Embedding: embedder,
	})
	if err != nil {
		panic(err)
	}

	results, err := retriever.Retrieve(ctx, "h")
	if err != nil {
		panic(err)
	}
	fmt.Println(results)
}
