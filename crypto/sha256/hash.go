package sha256

import "crypto/sha256"

func HashData(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}
