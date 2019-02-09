package main

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/riquellopes/golf/api"
	"github.com/riquellopes/golf/crawler"
)

func main() {
	fiis := make(chan []crawler.FII)
	go crawler.Do(fiis, new(crawler.FiiCollector))

	app := echo.New()
	app.SetDebug(true)

	fii := <-fiis
	app.GET("/", api.Index)
	app.GET("/api/", api.AllF(fii))
	app.GET("/api/:code", api.CodeF(fii))

	port := os.Getenv("PORT")
	app.Run(standard.New(":" + port))
}
