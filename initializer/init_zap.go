package initializer

import (
	"auto-manager/global"
	"auto-manager/utils"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var level zapcore.Level

func InitZap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.AM_CONFIG.Zap.Directory); !ok {
		fmt.Printf(" *** create directory *** \n%v\n", global.AM_CONFIG.Zap.Directory)
		_ = utils.MkDir(global.AM_CONFIG.Zap.Directory)
	}

	switch global.AM_CONFIG.Zap.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.PanicLevel
	default:
		level = zap.InfoLevel
	}

	if level == zap.DebugLevel || level == zap.ErrorLevel {
		logger = zap.New(getEncoderCore(), zap.AddStacktrace(level))
	} else {
		logger = zap.New(getEncoderCore())
	}
	if global.AM_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.AM_CONFIG.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}

	switch global.AM_CONFIG.Zap.EncodeLevel {
	case "LowercaseLevelEncoder":
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case "LowercaseColorLevelEncoder":
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case "CapitalLevelEncoder":
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case "CapitalColorLevelEncoder":
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

func getEncoder() zapcore.Encoder {
	if global.AM_CONFIG.Zap.Format == "json" {
		zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

func getEncoderCore() zapcore.Core {
	writer, err := utils.GetWriteSyncer()
	if err != nil {
		fmt.Printf(" *** Get Write Syncer Failed ***\n%s\n", err.Error())
		return nil
	}
	return zapcore.NewCore(getEncoder(), writer, level)
}

func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(global.AM_CONFIG.Zap.Prefix + " 2006-01-02 15:04:05"))
}
