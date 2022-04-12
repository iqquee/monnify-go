package helper

import "encoding/base64"

var (
	ServerErr = "error occured while sending request to the server"
)

func Base64(text string) string {
	encodedText := base64.StdEncoding.EncodeToString([]byte(text))
	return encodedText
}
