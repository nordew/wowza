package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"wowza/internal/cache"
	"wowza/internal/config"
	httpHandler "wowza/internal/handler/http"
	"wowza/internal/service"
	"wowza/internal/storage"

	"wowza/pkg/db/dragonfly"
	"wowza/pkg/db/minio"
	postgres "wowza/pkg/db/postgres"
	"wowza/pkg/generator"
	"wowza/pkg/hash"
	"wowza/pkg/logger"
	"wowza/pkg/paseto"

	"go.uber.org/zap"
)

func Run() {
	zapLogger, err := logger.New()
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	defer zapLogger.Sync()

	cfg, err := config.Load()
	if err != nil {
		zapLogger.Fatal("failed to load config", zap.Error(err))
	}

	pgsql, err := postgres.PostgresConnect(cfg.Postgres)
	if err != nil {
		zapLogger.Fatal("failed to connect to postgres", zap.Error(err))
	}

	dfly, err := dragonfly.DragonflyConnect(context.Background(), cfg.Dragonfly)
	if err != nil {
		zapLogger.Fatal("failed to connect to dragonfly", zap.Error(err))
	}

	minioClient, err := minio.New(cfg.Minio)
	if err != nil {
		zapLogger.Fatal("failed to connect to minio", zap.Error(err))
	}

	// Create unified storages
	storages := storage.NewStorages(storage.Dependencies{
		DB:          pgsql,
		MinioClient: minioClient,
		MinioConfig: cfg.Minio,
	})

	// Create external dependencies
	cache := cache.New(dfly)
	passwordHasher := hash.New()
	pasetoManager := paseto.NewManager([]byte(cfg.Paseto.SymmetricKey))
	generator := generator.New()

	// Create unified services
	services := service.NewServices(service.Dependencies{
		Storages:       storages,
		Logger:         zapLogger,
		PasswordHasher: passwordHasher,
		PasetoManager:  pasetoManager,
		Cache:          cache,
		Generator:      generator,
	})

	handler := httpHandler.NewHandler(
		zapLogger,
		services,
		cfg.App.DBTimeout,
	)
	server := handler.InitRoutes()

	go func() {
		addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
		if err := server.Listen(addr); err != nil && !errors.Is(err, http.ErrServerClosed) {
			zapLogger.Fatal("failed to start server", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	zapLogger.Info("shutting down server...")

	if err := server.ShutdownWithTimeout(5 * time.Second); err != nil {
		zapLogger.Fatal("server forced to shutdown", zap.Error(err))
	}

	zapLogger.Info("server exiting")
}
