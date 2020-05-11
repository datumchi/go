package identity

import (
	"github.com/datumchi/go/generated/protocol"
)


type IdentityAttributeVerificationResult struct {

	ValidAttestedAttributes map[string]*protocol.Identity_Attribute
	ValidAttestors map[string][]*protocol.Address

	InvalidAttestedAttributes map[string]*protocol.Identity_Attribute
	InvalidAttestors map[string][]*protocol.Address

}




