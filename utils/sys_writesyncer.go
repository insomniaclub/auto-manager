package utils

import (
	"auto-manager/global"
	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)

func GetWriteSyncer() (zapcore.WriteSyncer, error) {
	writer, err := zaprotatelogs.New(
		path.Join(global.AM_CONFIG.Zap.Directory, "%Y-%m-%d.log"),
		zaprotatelogs.WithLinkName(global.AM_CONFIG.Zap.LinkName),
		zaprotatelogs.WithMaxAge(7*24*time.Hour),
		zaprotatelogs.WithRotationTime(24*time.Hour),
	)
	if global.AM_CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(writer)), err
	}
	return zapcore.AddSync(writer), err
}
