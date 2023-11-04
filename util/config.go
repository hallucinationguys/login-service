package util

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBSource          string `mapstructure:"DB_SOURCE"`
	ClientOrigin      string `mapstructure:"CLIENT_ORIGIN"`
	HTTPServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`

	AccessTokenExpiresIn  time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRED_IN"`
	RefreshTokenExpiresIn time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRED_IN"`
	AccessTokenMaxAge     int           `mapstructure:"ACCESS_TOKEN_MAXAGE"`
	RefreshTokenMaxAge    int           `mapstructure:"REFRESH_TOKEN_MAXAGE"`

	SecretKey    string `mapstructure:"SecretKey"`
	RefreshToken string `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY"`
	MigrationURL string `mapstructure:"MIGRATION_URL"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigFile("production.env")
	//viper.SetConfigFile(".env")
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
