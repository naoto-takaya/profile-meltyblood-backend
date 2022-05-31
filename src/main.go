package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/naoto-takaya/profile-meltyblood-backend/configs"
	"github.com/naoto-takaya/profile-meltyblood-backend/handler"
)

func main() {
	e := echo.New()

	var config configs.Config
	err := envconfig.Process("", &config)
	if err != nil {
		e.Logger.Fatal(err.Error())
	}
	e.Validator = configs.NewValidator()
	e.Use(configs.CustomContextMiddleware)
	e.Use(configs.ConfigMiddleware(config))

	// ミドルウェアを設定
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルートを設定
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "ok")
	})
	e.POST("/upload", handler.Upload)

	// サーバーをポート番号1323で起動
	e.Logger.Fatal(e.Start(":80"))
}
