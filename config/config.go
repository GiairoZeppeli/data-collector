package config

import (
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
	kafkaAddressKey  = "kafka.address"
	kafkaTopicKey    = "kafka.topic"
)

type Settings struct {
	Mongo MongoSettings
	Redis RedisSettings
	Kafka KafkaSettings
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

type KafkaSettings struct {
	Address []string
	Topic   []string
}

func NewSettings() (Settings, error) {
	err := initConfig()
	if err != nil {
		return Settings{}, err
	}

	return Settings{
		Kafka: KafkaSettings{
			Address: viper.GetStringSlice(kafkaAddressKey),
			Topic:   viper.GetStringSlice(kafkaTopicKey),
		},
	}, nil
}

func initConfig() error {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	return viper.ReadInConfig()
}
