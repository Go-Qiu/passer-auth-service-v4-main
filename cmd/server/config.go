package main

import (
	"os"
	"passer-auth-service-v4/pkg/controllers"

	"github.com/joho/godotenv"
)

// config is a struct for storing the configuration settings.
type Config struct {
	SERVER_ADDRESS      string
	DSN                 DSNConfig
	DATE_CREATED_FORMAT string
	JWT                 controllers.JWTConfig
}

// JWTConfig is a struct for storing the JWT configuration settings.
type JWTConfig struct {
	ISSUER     string
	EXP_MIN    string
	SECRET_KEY string
}

// DSNConfig is a struct for storing the Data Source Name configuration settings
// for each database to be accessed.
type DSNConfig struct {
	AUTH string
}

// getConfigurations return the configuration setings in the .env file
// in the same directory of the calling application file.
func getConfigurations() *Config {
	err := godotenv.Load()
	if err != nil {
		return nil
	}

	c := Config{}
	c.SERVER_ADDRESS = os.Getenv("SERVER_ADDR")
	c.DATE_CREATED_FORMAT = os.Getenv("DATE_CREATED_FORMAT")
	c.DSN.AUTH = os.Getenv("DSN_AUTH")
	c.JWT.ISSUER = os.Getenv("JWT_ISSUER")
	c.JWT.EXP_MIN = os.Getenv("JWT_EXP_MINUTES")
	c.JWT.SECRET_KEY = os.Getenv("JWT_SECRET_KEY")

	return &c
}
