package encoding

import (
	"encoding/base64"
)

func Encode(data []byte) string {
	return base64.StdEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}



func Decode(data string) ([]byte, error) {
	return base64.StdEncoding.WithPadding(base64.NoPadding).DecodeString(data)
}