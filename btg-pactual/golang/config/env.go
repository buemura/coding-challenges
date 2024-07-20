package config

import (
	"os"

	"github.com/spf13/viper"
)

var (
	HTTP_PORT    string
	DATABASE_URL string
	BROKER_URL   string
)

func LoadEnv() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		HTTP_PORT = os.Getenv("HTTP_PORT")
		DATABASE_URL = os.Getenv("DATABASE_URL")
		BROKER_URL = os.Getenv("BROKER_URL")
		return
	}

	HTTP_PORT = viper.GetString("HTTP_PORT")
	DATABASE_URL = viper.GetString("DATABASE_URL")
	BROKER_URL = viper.GetString("BROKER_URL")
}
