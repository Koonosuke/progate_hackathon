package chatgptapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ChatGPTAPIInterface interface {
	SummarizeText(text string) (string, error)
}

type ChatGPTAPI struct{
	APIKey string
}

// CHatGPTAPIのコンストラクタ
func NewChatGPTAPI(apiKey string) ChatGPTAPIInterface {
	return &ChatGPTAPI{APIKey: apiKey}
}

// OpenAI APIにリクエストを送る
type ChatRequest struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
	MaxTokens int `json:"max_tokens"`
}

// OpenAI APIからのレスポンス
type ChatResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// テキストを要約したやつを取得する
func (api *ChatGPTAPI) SummarizeText(text string) (string, error) {
	prompt := fmt.Sprintf("テキストを要約してください:\n\n%s", text)

	reqPayload := ChatRequest{
		Model:     "gpt-3.5-turbo",
		MaxTokens: 500,
		Messages: []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{
			{Role: "system", Content: "You are a helpful assistant."},
			{Role: "user", Content: prompt},
		},
	}

	jsonData, err := json.Marshal(reqPayload)
	if err != nil {
		return "", fmt.Errorf("JSONエンコード失敗: %v", err)
	}

	apiURL := "https://api.openai.com/v1/chat/completions"
	apiKey := os.Getenv("OPENAI_APIKEY")
	if apiKey == "" {
		return "", fmt.Errorf("OPENAI_APIKEYが設定されていません")
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("リクエスト作成失敗: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("OpenAI APIリクエスト失敗: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("APIリクエスト失敗: ステータスコード: %d, レスポンス内容: %s", resp.StatusCode, body)
	}

	var respPayload ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&respPayload); err != nil {
		return "", fmt.Errorf("レスポンスデコード失敗: %v", err)
	}

	if len(respPayload.Choices) == 0 {
		return "", fmt.Errorf("要約を取得できませんでした")
	}

	return respPayload.Choices[0].Message.Content, nil
}