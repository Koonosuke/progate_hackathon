package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SlackUsecase interface {
	SendWeeklySummaryToSlack() error
}

type slackUsecase struct {
	db *pgxpool.Pool
}

func NewSlackUsecase(db *pgxpool.Pool) SlackUsecase {
	return &slackUsecase{db: db}
}

type stickyNote struct {
	Category string
	Content  string
}

func (s *slackUsecase) SendWeeklySummaryToSlack() error {
	query := `
		SELECT category, content FROM sticky
		WHERE created_at >= date_trunc('week', current_date)
		ORDER BY category, created_at
	`

	rows, err := s.db.Query(context.Background(), query)
	if err != nil {
		return err
	}

	defer rows.Close()

	var reflection []stickyNote
	for rows.Next() {
		var note stickyNote
		if err := rows.Scan(&note.Category, &note.Content); err != nil {
			return err
		}
		reflection = append(reflection, note)
	}

	// カテゴリ別でメッセージ構築
	var reflectionText, improvementText string
	for _, note := range reflection {
		if note.Category == "反省点" {
			reflectionText += fmt.Sprintf("- %s\n", note.Content)
		} else if note.Category == "改善点" {
			improvementText += fmt.Sprintf("- %s\n", note.Content)
		}
	}

	message := fmt.Sprintf(
		"📅【今週の振り返り】（%s〜）\n\n🔴 *反省点:*\n%s\n🟢 *改善点:*\n%s",
		time.Now().Format("2006-01-02"),
		reflectionText,
		improvementText,
	)

	payload := map[string]string{
		"text": message,
	}
	jsonData, _ := json.Marshal(payload)

	webhookURL := os.Getenv("SLACK_WEBHOOK_URL")
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("slack通知失敗 status: %d", resp.StatusCode)

	}
	return nil
}