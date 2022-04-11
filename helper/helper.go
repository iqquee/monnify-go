package helper

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnv() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	return nil
}

var Base_url = os.Getenv("MONNIFY_BASE_URL")
var Basic_Auth = os.Getenv("BASIC_AUTH")
var (
	EnvErr    = ".env file not found"
	ServerErr = "error occured while sending request to the server"
)
