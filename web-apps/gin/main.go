package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Ginエンジンのインスタンスを作成
	r := gin.Default()

	// ルートURL("/")に対するGETリクエストを処理するハンドラを設定
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	// 8080ポートでサーバーを起動
	r.Run() // デフォルトでは":8080"でリッスン
}
