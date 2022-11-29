package config

import (
	"github.com/spf13/viper"
	"log"
)

func NewConfig(fileName, directory string) (*Config, error) {
	var config Config

	if err := initConfig(fileName, directory); err != nil {
		log.Fatalf("crush init config, %v", err)
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &config, nil
}

func initConfig(fileName, directory string) error {
	viper.AddConfigPath(directory)
	viper.SetConfigName(fileName)
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
}
