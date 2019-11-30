package utils

import (
	"log"

	"go.uber.org/zap"
)

// Log returns a configured logger
func Log() *zap.SugaredLogger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Printf("could not init logger: %v", err)
	}
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	return sugar
}
