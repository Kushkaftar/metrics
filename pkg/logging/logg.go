package logging

import (
	"fmt"
	"log"
	"metrics/internal/models"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(c *models.Config) *zap.Logger {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(cfg)

	//todo: add check folder, and create folder
	path := fmt.Sprintf("%s/%s", c.Logs.Path, c.Logs.FileName)

	logFile, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	writer := zapcore.AddSync(logFile)
	defaultLogLevel := zapcore.InfoLevel
	core := zapcore.NewTee(zapcore.NewCore(fileEncoder, writer, defaultLogLevel))

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return logger
}
