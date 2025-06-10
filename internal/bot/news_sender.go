package bot

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/kopo-k/go-discord-newsbot/internal/news"
)

// SendNewsToDiscord は取得したニュースをDiscordに投稿します
func SendNewsToDiscord(s *discordgo.Session) error {
	channelID := os.Getenv("CHANNEL_ID")
	if channelID == "" {
		return fmt.Errorf("CHANNEL_ID が設定されていません")
	}

	articles, err := news.GetTopNews()
	if err != nil {
		return err
	}

	if len(articles) == 0 {
		s.ChannelMessageSend(channelID, "📰 今日のニュースは見つかりませんでした。")
		return nil
	}

	message := "📢 今日の注目ニュース（自動配信）\n"
	for _, article := range articles {
		message += fmt.Sprintf("• [%s](%s)\n", article.Title, article.URL)
	}

	_, err = s.ChannelMessageSend(channelID, message)
	return err
}
