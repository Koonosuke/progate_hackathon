package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/matthewTechCom/progate_hackathon/internal/usecase"
)

type SlackController struct {
	usecase usecase.SlackUsecase
}

func NewSlackController(e *echo.Echo, uc usecase.SlackUsecase) {
	h := &SlackController{usecase: uc}
	e.POST("/slack/weekly-summary", h.SendWeeklySummary)
}

func (h *SlackController) SendWeeklySummary(c echo.Context) error {
	err := h.usecase.SendWeeklySummaryToSlack()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Slack通知を送信しました"})
}
