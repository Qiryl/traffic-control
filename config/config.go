package config

import (
	"errors"
	"github.com/spf13/viper"
    "github.com/sirupsen/logrus"
)

type Config struct {
    Server ServerConfig
    Logger LoggerConfig
    File   FileConfig
}

type ServerConfig struct {
    Port string
}

type LoggerConfig struct {
    Level string
}

type FileConfig struct {
    Path string
}

func LoadConfig(filename string) (*viper.Viper, error) {
    v := viper.New()
    v.SetConfigName(filename)
    v.AddConfigPath(".")

    if err := v.ReadInConfig(); err != nil {
        if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
    }

    return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		logrus.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}

