package controller

import (
	"net/http"

	"github.com/matthewTechCom/progate_hackathon/internal/usecase"
	"github.com/matthewTechCom/progate_hackathon/internal/validator"
	"github.com/labstack/echo/v4"
)

type WidgetControllerInterface interface {
	ProcessBoard(ctx echo.Context) error
	GetAllStickies(ctx echo.Context) error
}

type WidgetRequest struct {
	BoardID     string `json:"boardID" validate:"required"`
	AccessToken string `json:"accessToken" validate:"required"`
}

type WidgetController struct {
	Usecase   usecase.WidgetUsecaseInterface
	Validator validator.ValidatorInterface
}

// コンストラクタ
func NewWidgetController(uc usecase.WidgetUsecaseInterface, v validator.ValidatorInterface) WidgetControllerInterface {
	return &WidgetController{
		Usecase:   uc,
		Validator: v,
	}
}

func (c *WidgetController) ProcessBoard(ctx echo.Context) error {
	req := new(WidgetRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "リクエストのフォーマットが不正です"})
	}
	if err := c.Validator.Validate(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "無効なリクエストパラメータ"})
	}

	// ユースケースの実行:miroのボードから情報を取得してPostgreSQLに保存する
	savedIDs, err := c.Usecase.ProcessAndSave(req.BoardID, req.AccessToken)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "情報をDBに保存しました",
		"ids": savedIDs,
	})
}
func (c *WidgetController) GetAllStickies(ctx echo.Context) error {
	stickies, err := c.Usecase.GetAllStickies()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, stickies)
}

