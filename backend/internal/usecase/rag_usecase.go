package usecase

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/matthewTechCom/progate_hackathon/internal/repository"
)

type RAGUsecase struct {
	Repo *repository.RAGRepository
}

type OpenAIResponse struct {
	Data []struct {
		Embedding []float64 `json:"embedding"`
	} `json:"data"`
}

// OpenAI の API を使って埋め込みを取得
func getEmbedding(query string) ([]float64, error) {
	apiKey := os.Getenv("OPENAI_APIKEY")
	if apiKey == "" {
		return nil, errors.New("OpenAI API key is missing")
	}

	body, _ := json.Marshal(map[string]interface{}{
		"input": query,
		"model": "text-embedding-ada-002",
	})

	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/embeddings", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res OpenAIResponse
	json.NewDecoder(resp.Body).Decode(&res)
	if len(res.Data) == 0 {
		return nil, errors.New("embedding is empty")
	}

	return res.Data[0].Embedding, nil
}

// クエリからカテゴリを抽出
func extractCategory(query string) (string, error) {
	if strings.Contains(query, "反省点") {
		log.Println("カテゴリ抽出: 反省点")
		return "反省点", nil
	} else if strings.Contains(query, "改善点") {
		log.Println("カテゴリ抽出: 改善点")
		return "改善点", nil
	}
	log.Println("カテゴリ抽出: なし")
	return "", errors.New("カテゴリが見つかりません")
}

// Embedding を更新するユースケース
func (u *RAGUsecase) EnsureEmbedding(id int, query string) error {
	log.Println("EnsureEmbedding called with ID:", id, "Query:", query)

	embedding, err := getEmbedding(query)
	if err != nil {
		log.Println("Failed to get embedding:", err)
		return err
	}
	log.Println("Embedding generated. Updating DB...")

	err = u.Repo.UpdateEmbedding(id, embedding)
	if err != nil {
		log.Println("Failed to update DB:", err)
	}
	return err
}

// 類似検索を行うユースケース（カテゴリ付き）
func (u *RAGUsecase) SearchWithRAG(query string) ([]map[string]string, error) {
	embedding, err := getEmbedding(query)
	if err != nil {
		log.Println("Failed to get embedding:", err)
		return nil, err
	}

	category, err := extractCategory(query)
	if err != nil {
		log.Println("Category not found, searching without category filter.")
		category = ""
	} else {
		log.Println("Category extracted:", category)
	}

	log.Println("Performing search with category:", category)
	results, err := u.Repo.SearchStickies(embedding, category)
	if err != nil {
		log.Println("Search failed:", err)
	}
	return results, err
}


func generateAdvice(query string, similarData []map[string]string) (string, error) {
	prompt := "以下のユーザーの悩みに対し、過去の改善点をもとに具体的なアドバイスを1つ出してください。\n\n"
	prompt += "【ユーザーの悩み】\n" + query + "\n\n"
	prompt += "【過去の改善点】\n"
	for _, item := range similarData {
		prompt += "- " + item["content"] + "\n"
	}
	prompt += "\n【アドバイス】"

	body, _ := json.Marshal(map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]string{
			{"role": "system", "content": "あなたは優秀なスクラムマスターです。"},
			{"role": "user", "content": prompt},
		},
	})

	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_APIKEY"))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var res struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	json.NewDecoder(resp.Body).Decode(&res)

	if len(res.Choices) == 0 {
		return "", errors.New("アドバイス生成に失敗しました")
	}

	return res.Choices[0].Message.Content, nil
}

func (u *RAGUsecase) SearchWithRAGAndAdvice(query string) (map[string]interface{}, error) {
	results, err := u.SearchWithRAG(query)
	if err != nil {
		return nil, err
	}

	advice, err := generateAdvice(query, results)
	if err != nil {
		advice = "アドバイスの生成に失敗しました。"
	}

	return map[string]interface{}{
		"results": results,
		"advice":  advice,
	}, nil
}

