package main

import (
	"fmt"
	"os"

	"github.com/kopo-k/go-discord-newsbot/config" // ←ここを自分のモジュール名に
)

func main() {
	config.LoadEnv()

	fmt.Println("DISCORD_TOKEN:", os.Getenv("DISCORD_TOKEN"))
	fmt.Println("NEWS_API_KEY:", os.Getenv("NEWS_API_KEY"))
	fmt.Println("CHANNEL_ID:", os.Getenv("CHANNEL_ID"))
}
