package config

import (
	"strings"

	"github.com/achintya-7/dating-app/logger"
	"github.com/spf13/viper"
)

var Values *Config

type Config struct {
	MySqlUrl          string `mapstructure:"MYSQL_URL"`
	RedisUrl          string `mapstructure:"REDIS_URL"`
	HttpPort          string `mapstructure:"HTTP_PORT"`
	TokenSymmetricKey string `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	EmailName         string `mapstructure:"EMAIL_NAME"`
	EmailAddress      string `mapstructure:"EMAIL_ADDRESS"`
	EmailPassowrd     string `mapstructure:"EMAIL_PASSWORD"`
}

// LoadConfig loads the configuration from the config file or environment variables
func LoadConfig() (*Config) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.AutomaticEnv()

	config := &Config{}

	err := viper.ReadInConfig()
	if err != nil {
		if strings.Contains(err.Error(), "Not Found") {
			logger.Warn(nil, "Config file not found, using environment variables")
			LoadEnvVariables(config)
		} else {
			logger.Fatal(nil, "Error reading config file or maybe it was not present: ", err.Error())
		}
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		logger.Fatal(nil, "Unable to decode into struct, ", err.Error())
	}

	logger.Info(nil, "Config loaded successfully")
	logger.Debug(nil, "Config: ", config)

	Values = config
	return config
}

func LoadEnvVariables(config *Config) {
	if config.MySqlUrl == "" {
		config.MySqlUrl = viper.GetString("MYSQL_URL")
	}
	if config.RedisUrl == "" {
		config.RedisUrl = viper.GetString("REDIS_URL")
	}
	if config.HttpPort == "" {
		config.HttpPort = viper.GetString("HTTP_PORT")
	}
	if config.TokenSymmetricKey == "" {
		config.TokenSymmetricKey = viper.GetString("TOKEN_SYMMETRIC_KEY")
	}
	if config.EmailName == "" {
		config.EmailName = viper.GetString("EMAIL_NAME")
	}
	if config.EmailAddress == "" {
		config.EmailAddress = viper.GetString("EMAIL_ADDRESS")
	}
	if config.EmailPassowrd == "" {
		config.EmailPassowrd = viper.GetString("EMAIL_PASSWORD")
	}
}
