package ed25519

import "golang.org/x/crypto/ed25519"

func Sign(privateKey []byte, data []byte) []byte {
	pk := ed25519.NewKeyFromSeed(privateKey)
	sig := ed25519.Sign(pk, data)

	return sig
}


func Verify(publicKey []byte, signature []byte, dataSigned []byte) bool {
	return ed25519.Verify(publicKey, dataSigned, signature)
}