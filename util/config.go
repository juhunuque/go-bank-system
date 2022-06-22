package util

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

// Config stores all configuration for the application
// The values are read by viper from a config file or environment variable
type Config struct {
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	DBSource            string        `mapstructure:"DB_SOURCE"`
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey   string        `mapstructure:TOKEN_SYMMETRIC_KEY`
	AccessTokenDuration time.Duration `mapstructure:ACCESS_TOKEN_DURATION`
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigFile(fmt.Sprintf("%s%s", path, ".env"))

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
