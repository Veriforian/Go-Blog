package main

import (
	"benjones/go-blog/handler"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	app.Static("/static", "static")

	// Handlers
	basicHandler := handler.BasicHandler{
		Name: "Ben",
	}

	app.GET("/", basicHandler.ShowHome)

	port := os.Getenv("PORT")

	if port == "" {
		port = ":3000"
	}

	app.Logger.Fatal(app.Start(port))
}
