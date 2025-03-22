package router

import (
	"github.com/labstack/echo/v4"
	"github.com/matthewTechCom/progate_hackathon/internal/controller"
)

func SetupRoutes(e *echo.Echo, boardController controller.BoardSummaryControllerInterface) {
	e.GET("/process-board", boardController.ProcessBoard)
}
