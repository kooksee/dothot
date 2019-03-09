package main

import (
	"github.com/kooksee/dothot/gens"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main() {

	// web
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.DefaultLoggerConfig))

	// static
	e.GET("/static/*", echo.WrapHandler(http.FileServer(gens.FS(false))))

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
