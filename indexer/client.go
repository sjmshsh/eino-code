package main

import (
	"context"
	"log"

	cli "github.com/milvus-io/milvus-sdk-go/v2/client"
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
