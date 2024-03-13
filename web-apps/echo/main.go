package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func main() {
	// Echoインスタンスの作成
	e := echo.New()

	// ルート(/)へのGETリクエストを処理するハンドラーを定義
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// サーバーを8080ポートで起動
	e.Start(":8080")
}

