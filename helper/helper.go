package helper

import (
	"os"
)

var Base_url = os.Getenv("MONNIFY_BASE_URL")
var Basic_Auth = os.Getenv("BASIC_AUTH")
var (
	EnvErr    = ".env file not found"
	ServerErr = "error occured while sending request to the server"
)
