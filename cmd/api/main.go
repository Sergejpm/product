package main

import (
	"context"
	"github.com/sergejpm/product/internal/config"
	"github.com/sergejpm/product/internal/domain/service/authorization"
	"github.com/sergejpm/product/internal/infra/log"
	"github.com/sergejpm/product/internal/infra/repository"
	"github.com/sergejpm/product/internal/server"
	"go.uber.org/zap/zapcore"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()

	defer func() {
		_ = log.Logger().Sync()
	}()

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg, err := config.Load()
	if err != nil {
		log.Logger().Fatalf("failed load config: %v", err)
	}

	logLevel := zapcore.ErrorLevel

	if errParse := logLevel.UnmarshalText([]byte(cfg.LogLevel)); errParse != nil {
		log.Logger().Errorf("unable to parse log level: %v", errParse)
	}

	log.SetLevel(logLevel)

	db, err := repository.Open(ctx, repository.Creds{
		ConnectionString: cfg.DBConnectionString,
		MaxIdleConns:     cfg.DBMaxIdleConnections,
		MaxOpenConns:     cfg.DBMaxOpenConnections,
	})

	if err != nil {
		log.Logger().Fatalf("unable to open database connection: %v", err)
	}

	defer func() {
		_ = db.Close()
	}()

	productServer := server.NewServer(db, cfg.TokenKey)
	authService := authorization.NewService(repository.NewTokenDBRepository(db), repository.NewUserDBRepository(db), cfg.TokenKey)

	err = runHTTPServer(ctx, cfg.HTTPPort, productServer, authService)
	if err != nil {
		log.Logger().Fatalf("unable to start http server: %v", err)
	}
}
