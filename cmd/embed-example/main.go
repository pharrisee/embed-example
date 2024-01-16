package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pharrisee/embed-example/web"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Filesystem: http.FS(web.StaticFS),
	}))
	e.GET("/", func(c echo.Context) error {
		return c.File("web/views/index.html")
	})
	e.Start(":2112")
}
