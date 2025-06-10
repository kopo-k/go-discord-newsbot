package main

import (
	"fmt"

	"github.com/kopo-k/go-discord-newsbot/config"
	"github.com/kopo-k/go-discord-newsbot/internal/news"
)

func main() {
	config.LoadEnv()

	articles, err := news.GetTopNews()
	if err != nil {
		panic(err)
	}

	for _, article := range articles {
		fmt.Println("📰", article.Title)
		fmt.Println(article.URL)
		fmt.Println()
	}
}
