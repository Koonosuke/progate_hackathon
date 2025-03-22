package controller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/matthewTechCom/progate_hackathon/internal/usecase"
	"github.com/matthewTechCom/progate_hackathon/util"
)

type IauthController interface {
	LoginHandler(c echo.Context) error
	CallbackHandler(c echo.Context) error
}

type authController struct {
	authUsecase usecase.IAuthUsecase
}

func NewAuthController(authUsecase usecase.IAuthUsecase) *authController {
	return &authController{
		authUsecase: authUsecase,
	}
}

func (ac *authController) LoginHandler(c echo.Context) error {
	url := util.GoogleOauthConfig.AuthCodeURL("state")
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func (ac *authController) CallbackHandler(c echo.Context) error {
	code := c.QueryParam("code")
	if code == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "codeが存在しません"})
	}

	user, err := ac.authUsecase.HandleGoogleCallback(c.Request().Context(), code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	redirectURL := fmt.Sprintf("%s?userId=%d", os.Getenv("FRONTEND_URL"), user.ID)
	return c.Redirect(http.StatusTemporaryRedirect, redirectURL)
}
