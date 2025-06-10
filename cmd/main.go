package main

import (
	"log"

	"github.com/kopo-k/go-discord-newsbot/config"
	"github.com/kopo-k/go-discord-newsbot/internal/bot"
)

func main() {
	config.LoadEnv()

	err := bot.StartBot()
	if err != nil {
		log.Fatalf("Bot 起動エラー: %v", err)
	}
}
