package middleware

import (
	"net/http"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

// CORS設定を行う
func CORSMiddleware() echo.MiddlewareFunc {
	return echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: []string{"*"}, // 必要に応じて許可するオリジンを指定
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization", "X-CSRF-Token"},
	})
}

// CSRF保護を行うミドルウェア
func CSRFMiddleware() echo.MiddlewareFunc {
	return echoMiddleware.CSRFWithConfig(echoMiddleware.CSRFConfig{
		TokenLookup: "header:X-CSRF-Token", // CSRFトークンはヘッダーから取得する
		CookieName: "csrf",
		ErrorHandler: func(err error, c echo.Context) error {
			return c.JSON(http.StatusForbidden, map[string]string{
				"error": "CSRFトークンが無効です",
			})
		},
	})
}

