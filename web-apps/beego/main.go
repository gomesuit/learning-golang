package main

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func main() {
	// ルートURL("/")に対するGETリクエストを処理する関数を定義
	web.Get("/", func(ctx *context.Context) {
		ctx.WriteString("Hello, Beego!")
	})

	// Webサーバーを8080ポートで起動
	web.Run()
}
