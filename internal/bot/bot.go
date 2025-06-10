package bot

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Botを起動する関数
// func StartBot() error {
// 	token := os.Getenv("DISCORD_TOKEN")
// 	if token == "" {
// 		return fmt.Errorf("DISCORD_TOKEN が設定されていません")
// 	}

// 	// Discordセッション作成
// 	dg, err := discordgo.New("Bot " + token)
// 	if err != nil {
// 		return fmt.Errorf("Discordセッションの作成に失敗: %v", err)
// 	}

// 	// イベントを追加（Botが起動したときにコンソール出力）
// 	dg.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
// 		fmt.Println("✅ Bot is now connected!")
// 	})

// 	// 接続開始
// 	err = dg.Open()
// 	if err != nil {
// 		return fmt.Errorf("Discordに接続できませんでした: %v", err)
// 	}

// 	// CTRL+C で停止するまで待つ
// 	fmt.Println("🤖 Bot is running. Press CTRL+C to exit.")
// 	stop := make(chan os.Signal, 1)
// 	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
// 	<-stop

// 	// 終了処理
// 	dg.Close()
// 	fmt.Println("🛑 Bot is shutting down.")
// 	return nil
// }

func StartBot() error {
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		return fmt.Errorf("DISCORD_TOKEN が設定されていません")
	}

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return fmt.Errorf("Discordセッション作成失敗: %v", err)
	}

	// Botがオンラインになったときのメッセージ
	dg.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Println("Bot is now running as", s.State.User.Username)
	})

	// メッセージへの応答ハンドラ
	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.Bot {
			return
		}
		if m.Content == "!ping" {
			s.ChannelMessageSend(m.ChannelID, "Pong!")
		}
	})

	err = dg.Open()
	if err != nil {
		return fmt.Errorf("Botの接続に失敗: %v", err)
	}

	fmt.Println("Bot is running. Press Ctrl+C to exit.")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	dg.Close()
	return nil
}
