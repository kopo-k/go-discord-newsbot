package bot

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron/v3"
)

// StartScheduler は毎朝8時にニュースを送信するスケジューラーを開始します
func StartScheduler(session *discordgo.Session) {
	c := cron.New(cron.WithLocation(time.FixedZone("Asia/Tokyo", 9*60*60))) // JST固定
	_, err := c.AddFunc("0 8 * * *", func() {
		log.Println("⏰ 8時になったのでニュース送信します！")
		err := SendNewsToDiscord(session)
		if err != nil {
			log.Printf("❌ ニュース送信に失敗: %v\n", err)
		}
	})
	if err != nil {
		log.Fatalf("cron登録失敗: %v", err)
	}
	c.Start()
}
