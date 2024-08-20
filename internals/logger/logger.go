package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	Log     *zap.Logger
	logFile *os.File
)

func InitLogger(logLevel string) {
	config := zap.NewProductionConfig()

	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var err error
	logFile, err = os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("Failed to open log file: " + err.Error())
	}
	config.OutputPaths = []string{"stdout", logFile.Name()}

	switch logLevel {
	case "debug":
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		config.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		config.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	default:
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	Log, err = config.Build()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
}

func Close() {
	if Log != nil {
		_ = Log.Sync() // flush any buffered log entries
		Log = nil      // stop logging
	}
	if logFile != nil {
		_ = logFile.Close() // close the log file
		logFile = nil       // ensure the log file is not used again
	}
}
