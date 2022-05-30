package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/spf13/pflag"
	"go.uber.org/zap"

	"github.com/bekryasheva/url-shortener/internal/app"
	"github.com/bekryasheva/url-shortener/internal/app/handlers"
	"github.com/bekryasheva/url-shortener/internal/app/storage"
	"github.com/bekryasheva/url-shortener/internal/app/storage/local"
	"github.com/bekryasheva/url-shortener/internal/app/storage/postgres"
)

const (
	DefaultConfigPath = "configs/config.yaml"
	PostgresqlStorage = "postgresql"
	LocalStorage      = "local"
)

func main() {
	log, err := zap.NewProduction()
	if err != nil {
		fmt.Printf("failed to initialize logger: %v\n", err)
		return
	}
	defer log.Sync()

	configFile := pflag.StringP("config", "c", DefaultConfigPath, "Setting the configuration file")

	pflag.Parse()

	cfg, err := app.ReadConfigFromFile(*configFile)
	if err != nil {
		log.Fatal("failed to read config", zap.Error(err))
	}

	log.Info("using config file", zap.String("config_file", *configFile))

	var urlStorage storage.Storage

	switch cfg.Storage {
	case PostgresqlStorage:
		db, err := postgres.NewPostgresDB(cfg, log)
		if err != nil {
			log.Fatal("failed to open a DB connection", zap.Error(err))
		}
		defer db.Close()

		urlStorage = postgres.NewPostgresDatabase(db)
	case LocalStorage:
		urlStorage = local.NewLocalStorage()
	}

	e := echo.New()

	e.POST("/url", handlers.SaveHandler(urlStorage, cfg.URL.AddressPrefix))
	e.GET("/url/:url", handlers.GetHandler(urlStorage))
	e.GET("/:url", handlers.RedirectHandler(urlStorage))

	e.Logger.Fatal(e.Start(cfg.API.Address))
}
