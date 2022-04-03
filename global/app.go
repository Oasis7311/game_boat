package global

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"

	"github.com/oasis/game_boat/config"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
	Log         *zap.Logger
	DB          *gorm.DB
	Redis       *redis.Client
}

var App = new(Application)
