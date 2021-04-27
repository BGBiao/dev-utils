package encoding

import (
	"encoding/base64"
)

func Base64(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func DecodeBase64(data string) ([]byte, error) {

	return base64.StdEncoding.DecodeString(data)
}
