package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load() // 加载环境变量
	if err != nil {
		log.Fatal("Error loading .env file") // 处理加载错误
	}

	ctx := context.Background()
	model, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: os.Getenv("ARK_API_KEY"),
		Model:  os.Getenv("MODEL"),
	})
	//input := []*schema.Message{
	//	schema.SystemMessage("你是一个可爱的高中美少女"),
	//	schema.UserMessage("你好"),
	//}
	//response, err := model.Generate(ctx, input)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(response.Content)

	//reader, err := model.Stream(ctx, input)
	//if err != nil {
	//	panic(err)
	//}
	//defer reader.Close()
	//
	//// 处理流式内容
	//for {
	//	chunk, err := reader.Recv()
	//	if err != nil {
	//		break
	//	}
	//	fmt.Print(chunk.Content)
	//}

	template := prompt.FromMessages(schema.FString,
		schema.SystemMessage("你是一个{role}"),
		&schema.Message{
			Role:    schema.User,
			Content: "请帮帮我，史瓦罗先生，{task}",
		},
	)
	params := map[string]any{
		"role": "机器人史瓦罗先生",
		"task": "写一首诗",
	}
	messages, err := template.Format(ctx, params)
	msg, err := model.Generate(ctx, messages)
	if err != nil {
		panic(err)
	}
	fmt.Print(msg)
}
