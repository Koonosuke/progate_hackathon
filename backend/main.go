package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/sashabaranov/go-openai"

	"github.com/matthewTechCom/progate_hackathon/internal/controller"
	"github.com/matthewTechCom/progate_hackathon/internal/infrastructure"
	"github.com/matthewTechCom/progate_hackathon/internal/miroapi"
	"github.com/matthewTechCom/progate_hackathon/internal/repository"
	"github.com/matthewTechCom/progate_hackathon/internal/usecase"
	"github.com/matthewTechCom/progate_hackathon/internal/validator"
)

func main() {
	fmt.Println("SLACK_WEBHOOK_URL from env:", os.Getenv("SLACK_WEBHOOK_URL"))

	db, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
	}))

	ragRepo := &repository.RAGRepository{DB: db}
	ragUsecase := &usecase.RAGUsecase{Repo: ragRepo}
	ragController := &controller.RAGController{Usecase: ragUsecase}
	e.GET("/update_embedding", ragController.UpdateEmbedding)
	e.GET("/search", ragController.SearchWithRAG)
	e.GET("/search_with_advice", ragController.SearchWithAdvice)
	




	slackUsecase := usecase.NewSlackUsecase(db)
	controller.NewSlackController(e, slackUsecase)
	e.POST("/slack/weekly-summary", func(c echo.Context) error {
		if err := slackUsecase.SendWeeklySummaryToSlack(); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, echo.Map{"message": "Slack通知を送信しました"})
	})

	boardRepo := repository.NewBoardRepository(db)
	stickyRepo := repository.NewStickyRepository(db)
	miroClient := miroapi.NewMiroAPI(os.Getenv("MIRO_ACCESS_TOKEN"))

	openaiClient := openai.NewClient(os.Getenv("OPENAI_APIKEY"))
	embedder := infrastructure.NewOpenAIEmbedding(openaiClient)

	widgetUsecase := usecase.NewWidgetUsecase(boardRepo, stickyRepo, miroClient, embedder)
	v := validator.NewValidator()
	widgetController := controller.NewWidgetController(widgetUsecase, v)

	e.POST("/process-board", widgetController.ProcessBoard)
	e.GET("/stickies", widgetController.GetAllStickies)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
