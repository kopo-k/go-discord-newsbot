package news

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const newsAPIURL = "https://newsapi.org/v2/top-headlines?country=us&category=technology&pageSize=5"


type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

type NewsResponse struct {
	Status  string    `json:"status"`
	Total   int       `json:"totalResults"`
	Articles []Article `json:"articles"`
}

// GetTopNews は最新のニュースを取得します
func GetTopNews() ([]Article, error) {
	apiKey := os.Getenv("NEWS_API_KEY")
	if apiKey == "" {
		fmt.Println("❌ 環境変数 NEWS_API_KEY が設定されていません")
		return nil, fmt.Errorf("NEWS_API_KEY が設定されていません")
	}

	fmt.Println("✅ APIキー取得済み")

	req, err := http.NewRequest("GET", newsAPIURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Println("✅ APIレスポンスステータス:", resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("APIエラー: ステータスコード %d", resp.StatusCode)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	var news NewsResponse
	if err := json.Unmarshal(body, &news); err != nil {
		return nil, err
	}

	fmt.Println("✅ ニュース件数:", len(news.Articles))

	return news.Articles, nil
}
