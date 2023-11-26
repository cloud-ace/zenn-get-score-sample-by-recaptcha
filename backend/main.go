package main

import (
	"net/http"
	"os"
	"re_capthcha_enterprise/backend/re_capcha"

	"github.com/labstack/echo/v4/middleware"

	echo "github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// CORSミドルウェアを追加
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // すべてのオリジンからのリクエストを許可
		AllowHeaders: []string{"*"}, // すべてのヘッダーを許可
	}))

	projectId := os.Getenv("PROJECT_ID")
	siteKey := os.Getenv("SITE_KEY")

	outputPath := os.Getenv("LOCAL_FILE_PATH")

	reCaptchaClient := re_capcha.NewClient(projectId, siteKey, outputPath)

	// ログイン処理エンドポイント
	e.POST("/login", func(c echo.Context) error {

		token := c.Request().Header.Get("re-captcha-token")
		action := c.Request().Header.Get("re-captcha-action")

		ctx := c.Request().Context()

		// reCAPTCHAによる評価を取得
		risk, err := reCaptchaClient.CreateAssessment(ctx, token, action)

		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		if risk.Score < 0.5 {
			// スコアが0.5未満の場合は、MFAを要求するなど再認証を要求するなどの対応を行う。
			// デモのため、今回はエラーを返す
			return c.String(http.StatusUnauthorized, "Invalid reCAPTCHA score")
		}

		// TODO: ログイン処理を実行

		return c.JSON(http.StatusOK, "success: login")
	})

	// サーバーを開始
	e.Start(":8081")
}
