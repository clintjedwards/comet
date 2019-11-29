package utils

import (
	"log"

	"go.uber.org/zap"
)

// Log returns a configured logger
// It configures itself based on appconfig
func Log() *zap.SugaredLogger {
	// config, err := config.FromEnv()
	// if err != nil {
	// 	return nil, err
	// }

	logger, err := zap.NewProduction()
	if err != nil {
		log.Printf("could not init logger: %v", err)
	}
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	return sugar
}
