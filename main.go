package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"gogolook/assets"
	"gogolook/di"
	myHttp "gogolook/http"
	"net/http"
	"os"
)

func main() {
	initializeDI := di.InitializeDI()
	go startHttpServer(initializeDI)
	select {}

}
func startHttpServer(di di.DI) {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURI:    true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Printf("REQUEST: uri: %v, status: %v\n", v.URI, v.Status)
			return nil
		},
	}))
	e.Use(middleware.Recover())

	// health check
	e.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	e.FileFS("/api-docs", "index.html", echo.MustSubFS(assets.IndexHTML, "swagger"))
	e.StaticFS("/api-docs", echo.MustSubFS(assets.Dist, "swagger"))

	myHttp.RegisterHandlers(e.Group(""), di.HttpServer)
	log.Fatal(e.Start(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))))
}
