package domain

import (
	"github.com/datumchi/go/crypto/ed25519"
	"github.com/datumchi/go/crypto/sha512"
	"github.com/datumchi/go/encoding"
	"github.com/datumchi/go/naming"
)

func VerifyDomainDefinition(definition naming.DomainDefinition) bool {

	// decode pubkey and signature
	pubkey, err := encoding.Decode(definition.PublicKey)
	if err != nil {
		return false
	}

	signature, err := encoding.Decode(definition.Signature)
	if err != nil {
		return false
	}

	hashedDomain := sha512.HashData([]byte(definition.Domain))
	return ed25519.Verify(pubkey, signature, hashedDomain)

}
