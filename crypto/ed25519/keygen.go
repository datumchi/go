package ed25519

import (
	ed "golang.org/x/crypto/ed25519"
	"io"
)

func GenerateKeypair(rand io.Reader) (publicKey []byte, privateKey []byte) {
	pub, pri, _ := ed.GenerateKey(rand)
	publicKey = pub
	privateKey = pri.Seed()

	return publicKey, privateKey

}



