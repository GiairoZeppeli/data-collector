package config

import (
	"github.com/spf13/viper"
)

const (
	configPath      = "config"
	configName      = "config.local"
	kafkaAddressKey = "kafka.address"
	kafkaTopicKey   = "kafka.topic"
)

type Settings struct {
	Kafka KafkaSettings
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
