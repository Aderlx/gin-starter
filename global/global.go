package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	// config
	Gol_CONFIG *viper.Viper
	// logger
	Gol_LOG *zap.Logger
	// db
	Gol_DB *gorm.DB
)
