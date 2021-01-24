package global

import (
	"auto-manager/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	AM_CONFIG *config.Server
	AM_VP     *viper.Viper
	AM_DB     *gorm.DB
	AM_LOG    *zap.Logger
)
