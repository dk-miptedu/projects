package configs

import (
	"FinalTaskAppGoBasic/logs"

	"github.com/spf13/viper"
)

type Config struct {
	DataBase Database `yaml:"database"`
	Restapi  Restapi  `yaml:"restapi"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

type Restapi struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Sslmode string `yaml:"sslmode"`
}

func LoadConfig(configFilePath string) (*Config, error) {
	viper.SetConfigFile(configFilePath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	settings := viper.AllSettings()
	logs.LogConfigsParams("config", settings)

	return &config, nil
}
