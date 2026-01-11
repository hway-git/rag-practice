package main

import (
	"fmt"
	"os"

	"rag-service/internal/config"
	"rag-service/pkg/logger"

	"go.uber.org/zap"
)

func main() {
	// 1. Load Configuration
	cfg, err := config.LoadConfig("configs")
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		os.Exit(1)
	}

	// 2. Initialize Logger
	if err := logger.Init(cfg.Log.Level, cfg.Server.Mode); err != nil {
		fmt.Printf("Failed to init logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.Sync()

	logger.Info("RAG Service Initializing...",
		zap.String("version", "v0.1.0"),
		zap.String("env", cfg.Server.Mode),
	)

	logger.Info("Configuration loaded successfully",
		zap.Int("port", cfg.Server.Port),
		zap.String("db_driver", cfg.Database.Driver),
	)

	// TODO: Init database connections
	// TODO: Init vector store client
	// TODO: Setup router and middleware
	// TODO: Start server

	logger.Info("Service initialization finished (Core components ready).")
}
