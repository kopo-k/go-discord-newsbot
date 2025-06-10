package bot

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func StartBot() error {
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		return fmt.Errorf("DISCORD_TOKEN が設定されていません")
	}

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return fmt.Errorf("Discordセッション作成失敗: %v", err)
	}

	// 起動メッセージ
	dg.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Println("Bot is now running as", s.State.User.Username)
	})

	// ✅ ← ここがあなたのコード部分（イベントハンドラ）
	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.Bot {
			return
		}

		switch m.Content {
		case "!ping":
			s.ChannelMessageSend(m.ChannelID, "Pong!")
		case "!news":
			err := SendNewsToDiscord(s)
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, "ニュース送信エラー: "+err.Error())
			}
		}
	})

	err = dg.Open()
	if err != nil {
		return fmt.Errorf("Botの接続に失敗: %v", err)
	}

	StartScheduler(dg)

	fmt.Println("Bot is running. Press Ctrl+C to exit.")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	dg.Close()
	return nil
}
