package main

import (
	"context"
	"log"
	"os"

	"github.com/matthewTechCom/progate_hackathon/internal/controller"
	"github.com/matthewTechCom/progate_hackathon/internal/repository"
	"github.com/matthewTechCom/progate_hackathon/internal/usecase"

	"github.com/joho/godotenv"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

func main() {
	godotenv.Load(".env")

	db, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	repo := &repository.RAGRepository{DB: db}
	usecase := &usecase.RAGUsecase{Repo: repo}
	ragController := &controller.RAGController{Usecase: usecase}

	e := echo.New()

	// エンドポイントの設定
	e.GET("/update_embedding", ragController.UpdateEmbedding)
	e.GET("/search", ragController.SearchWithRAG)

	e.Logger.Fatal(e.Start(":8080"))
}
