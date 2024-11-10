package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"gogolook/assets"
	"gogolook/di"
	myHttp "gogolook/http"
	"gogolook/lib/pg"
	"net/http"
	"os"
	"time"
)

func main() {
	migrateDB()
	initializeDI := di.InitializeDI()
	go startHttpServer(initializeDI)
	select {}

}
func migrateDB() {
	dbUrl := pg.GetPGURL()
	for {
		if err := isPGReady(dbUrl); err != nil {
			log.Warningln("PG is not ready, retry after 5 seconds")
			time.Sleep(5 * time.Second)
		}
		break
	}
	if err := migrateToLatest(dbUrl); err != nil {
		log.Fatal(err)
	}
}
func isPGReady(dbUrl string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return fmt.Errorf("sql.Open: %v", err)
	}
	defer db.Close()

	query := "select 1;"
	if _, err := db.ExecContext(ctx, query); err != nil {
		return fmt.Errorf("db.ExecContext: fail to connect to db, query: %v: %v", query, err)
	}

	return nil
}

func migrateToLatest(dbUrl string) error {
	m, err := migrate.New("file://migrations", dbUrl)
	if err != nil {
		return err
	}

	err = m.Up()
	version, _, _ := m.Version()

	if err == nil {
		log.WithField("version", version).Info("Migrated to latest version")
	} else if errors.Is(err, migrate.ErrNoChange) {
		log.WithField("version", version).Info("Already up-to-date")
	} else {
		return err
	}

	return nil
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
