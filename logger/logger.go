package logger

import (
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger = logr.Logger

const (
	DebugLevel int = -int(zapcore.DebugLevel)
	InfoLevel      = -int(zapcore.InfoLevel)
	WarnLevel      = -int(zapcore.WarnLevel)
	ErrorLevel     = -int(zapcore.ErrorLevel)
	FatalLevel     = -int(zapcore.FatalLevel)
)

var logDefault = func() Logger {
	config := zap.NewDevelopmentConfig()
	if os.Getenv("DEPLOYMENT") == "production" {
		config = zap.NewProductionConfig()
	}
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	zapLog, err := config.Build()
	if err != nil {
		panic(err)
	}

	zapr := zapr.NewLogger(zapLog).WithName("default")
	return zapr
}()

func LOG() Logger {
	return logDefault
}

func DEBUG() Logger {
	return logDefault.V(DebugLevel)
}

func WARNING() Logger {
	return logDefault.V(WarnLevel)
}
