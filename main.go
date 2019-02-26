package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	// web
	e := echo.New()

	// static
	e.GET("/static/*", echo.WrapHandler(http.FileServer(FS(false))))

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
