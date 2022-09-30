package main

import (
	"github.com/jankremlacek/go-bazel/shared"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info("servicea says:",
		zap.Int("sum", shared.Sum(1, 2)),
		zap.String("version", shared.Version()))
}