package config

import (
	"fmt"
	"github.com/spf13/viper"
)

const (
	configPath       = "config"
	configName       = "config.local"
	mongoDatabaseKey = "mongo.database"
	mongoPortKey     = "mongo.port"
	mongoHostKey     = "mongo.host"
	mongoUserKey     = "mongo.user"
	mongoPasswordKey = "mongo.password"
	redisPortKey     = "redis.port"
	redisHostKey     = "redis.host"
	redisUserKey     = "redis.user"
	redisPasswordKey = "redis.password"
	redisDBKey       = "redis.db"
)

type Settings struct {
	Mongo MongoSettings
	Redis RedisSettings
}

type MongoSettings struct {
	Database string
	User     string
	Password string
	MongoURL string
}

type RedisSettings struct {
	Address  string
	User     string
	Password string
	DB       int
}

func NewSettings() (Settings, error) {
	err := initConfig()
	if err != nil {
		return Settings{}, err
	}

	return Settings{
		Mongo: MongoSettings{
			Database: viper.GetString(mongoDatabaseKey),
			User:     viper.GetString(mongoUserKey),
			Password: viper.GetString(mongoPasswordKey),
			MongoURL: fmt.Sprintf("mongodb://%s:%s@%s:%s",
				viper.GetString(mongoUserKey),
				viper.GetString(mongoPasswordKey),
				viper.GetString(mongoHostKey),
				viper.GetString(mongoPortKey),
			),
		},
		Redis: RedisSettings{
			Address:  fmt.Sprintf("%s:%s", viper.GetString(redisHostKey), viper.GetString(redisPortKey)),
			User:     viper.GetString(redisUserKey),
			Password: viper.GetString(redisPasswordKey),
			DB:       viper.GetInt(redisDBKey),
		},
	}, nil
}

func initConfig() error {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	return viper.ReadInConfig()
}
