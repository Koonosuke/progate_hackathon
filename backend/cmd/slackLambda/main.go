package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	"github.com/matthewTechCom/progate_hackathon/internal/usecase"
)

func handler() error {
	_ = godotenv.Load() // .env ローカル開発用

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return fmt.Errorf("DATABASE_URL not set")
	}

	db, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		return fmt.Errorf("DB接続エラー: %v", err)
	}
	defer db.Close()

	slackUC := usecase.NewSlackUsecase(db)
	if err := slackUC.SendWeeklySummaryToSlack(); err != nil {
		return fmt.Errorf("slack送信エラー: %v", err)
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
