package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/eino-ext/components/embedding/ark"
	"github.com/cloudwego/eino-ext/components/indexer/milvus"
	"github.com/cloudwego/eino/schema"
	"github.com/joho/godotenv"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
	"log"
	"os"
	"time"
)

var collection = "test"

var fields = []*entity.Field{
	{
		Name:     "id",
		DataType: entity.FieldTypeVarChar,
		TypeParams: map[string]string{
			"max_length": "255",
		},
		PrimaryKey: true,
	},
	{
		Name:     "vector", // 确保字段名匹配
		DataType: entity.FieldTypeBinaryVector,
		TypeParams: map[string]string{
			"dim": "81920",
		},
	},
	{
		Name:     "content",
		DataType: entity.FieldTypeVarChar,
		TypeParams: map[string]string{
			"max_length": "8192",
		},
	},
	{
		Name:     "metadata",
		DataType: entity.FieldTypeJSON,
	},
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

	indexer, err := milvus.NewIndexer(ctx, &milvus.IndexerConfig{
		Client:     MilvusCli,
		Collection: collection,
		Fields:     fields,
		Embedding:  embedder,
	})
	if err != nil {
		log.Fatalf("Failed to create indexer: %v", err)
	}

	docs := []*schema.Document{
		{
			ID:      "1",
			Content: "你说得对，但是原神是一款二次元开放大世界游戏",
			MetaData: map[string]any{
				"author": "sunnyxyli",
			},
		},
		{
			ID:      "2",
			Content: "我喜欢你，hfj",
			MetaData: map[string]any{
				"author": "sunnyxyli",
			},
		},
	}
	ids, err := indexer.Store(ctx, docs)
	if err != nil {
		panic(err)
	}
	fmt.Println(ids)
}
