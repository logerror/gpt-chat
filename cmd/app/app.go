package app

import (
	"context"
	"fmt"
	"github.com/logerror/gpt-chat/configs"
	gogpt "github.com/sashabaranov/go-gpt3"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"strings"
)

var log *zap.Logger

func Run() error {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.AddConfigPath("$HOME")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatal("viper read in config failed", zap.Error(err))
	}
	var cfg configs.Configuration
	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}

	logger.Debug("config info", zap.Any("cfg", cfg))
	if len(cfg.OpenAISecretKey) == 0 {
		if len(os.Args) == 2 {
			cfg.OpenAISecretKey = os.Args[1]
		}
	}

	if len(cfg.OpenAISecretKey) == 0 {
		fmt.Println("You need to declare your secretkey in the following two ways")
		fmt.Println("1. gpt-chat you secret key")
		fmt.Println("2. export OPEN_AI_SECRET_KEY=\"you secret key\" && gpt-chat")
		os.Exit(1)
	}

	var request gogpt.ChatCompletionRequest
	client := gogpt.NewClient(cfg.OpenAISecretKey)

	ctx := context.Background()
	if cfg.Model == gogpt.GPT3Dot5Turbo0301 || cfg.Model == gogpt.GPT3Dot5Turbo {
		fmt.Print("input your question:")
		var messages []gogpt.ChatCompletionMessage
		for {
			var question string
			fmt.Scanln(&question)
			currentMessage := gogpt.ChatCompletionMessage{
				Role:    "user",
				Content: question,
			}
			request.Messages = append(messages, currentMessage)
			request.Model = cfg.Model
			resp, err := client.CreateChatCompletion(ctx, request)
			if err != nil {
				return err
			}
			fmt.Println(resp.Choices[0].Message.Content)
			messages = append(messages, gogpt.ChatCompletionMessage{
				Role:    "assistant",
				Content: resp.Choices[0].Message.Content,
			})
		}
	}

	return nil
}
