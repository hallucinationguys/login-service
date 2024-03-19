package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBSource          string `mapstructure:"DB_SOURCE"`
	ClientOrigin      string `mapstructure:"CLIENT_ORIGIN"`
	HTTPServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`

	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	AccessTokenMaxAge   int           `mapstructure:"ACCESS_TOKEN_MAXAGE"`

	SecretKey    string `mapstructure:"SecretKey"`
	MigrationURL string `mapstructure:"MIGRATION_URL"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	//viper.SetConfigFile("production.env")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	// Viper unmarshals the loaded env varialbes into the struct
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	return
}
