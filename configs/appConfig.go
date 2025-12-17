package configs

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
	Dsn string 
	AppSecret string
}

func SetupEnv() (cfg AppConfig, err error) {

if os.Getenv("APP_ENV") == "dev"{
	godotenv.Load()
}

	httpPort := os.Getenv("HTTP_PORT")
	if len(httpPort) <1 {
		return AppConfig{},errors.New("HTTP_PORT not set in environment")
	}

	Dsn:= os.Getenv("DSN")
	AppSecret:= os.Getenv("APP_SECRET")

	if len(Dsn) <1{
		return AppConfig{},errors.New("env variable not found")
	}

	return AppConfig{ServerPort: httpPort , Dsn :Dsn ,AppSecret: AppSecret} , nil
}