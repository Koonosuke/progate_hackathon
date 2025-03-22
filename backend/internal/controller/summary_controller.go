package controller

import (
	"net/http"
	"github.com/matthewTechCom/progate_hackathon/internal/usecase"
	"github.com/matthewTechCom/progate_hackathon/internal/validator"
	"github.com/labstack/echo/v4"
)

type BoardSummaryControllerInterface interface {
	ProcessBoard(ctx echo.Context) error
}

// リクエストの構造体（バリデーション用）
type BoardSummaryRequest struct {
	BoardID     string `json:"boardID" validate:"required"`
	AccessToken string `json:"accessToken" validate:"required"`
}

// controllerはユースケース、バリデータに依存する
type BoardSummaryController struct {
	Usecase usecase.BoardSummaryUsecaseInterface
	Validator validator.ValidatorInterface
}

// コンストラクタ
func NewBoardSummaryController(uc usecase.BoardSummaryUsecaseInterface, v validator.ValidatorInterface) BoardSummaryControllerInterface {
	return &BoardSummaryController{
		Usecase: uc,
		Validator: v,
	}
}

// ボードの要約を処理する
func (c *BoardSummaryController) ProcessBoard(ctx echo.Context) error {
	req := new(BoardSummaryRequest)

	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "リクエストのフォーマットが不正です"})
	}

	// バリデーション適用
	if err := c.Validator.Validate(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "無効なリクエストパラメータ"})
	}

	// ユースケースの実行
	savedIDs, err := c.Usecase.ProcessAndSave(req.BoardID, req.AccessToken)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "要約をDBに保存できました",
		"ids": savedIDs,
	})
}