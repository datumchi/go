package sha512

import (
	"crypto/sha512"
)

func HashData(data []byte) []byte {
	hash := sha512.Sum512(data)
	return hash[:]
}
