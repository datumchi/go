package identity

import (
	"context"
	"fmt"
	"github.com/datumchi/go/crypto/ed25519"
	"github.com/datumchi/go/crypto/sha256"
	"github.com/datumchi/go/encoding"
	"github.com/datumchi/go/generated/protocol"
	"github.com/datumchi/go/naming/dns"
	"github.com/datumchi/go/verifier/domain"
)


type IdentityAttributeVerificationResult struct {

	ValidAttestedAttributes map[string]*protocol.Identity_Attribute
	ValidAttestors map[string][]*protocol.Address

	InvalidAttestedAttributes map[string]*protocol.Identity_Attribute
	InvalidAttestors map[string][]*protocol.Address

}


func VerifyIdentity(ident protocol.Identity) bool {

	// This verifies domain
	domainDefinition, err := dns.GetDomainDefinition(ident.Address.Domain)
	if err != nil {
		return false
	}
	domainOk := domain.VerifyDomainDefinition(domainDefinition)
	if !domainOk {
		return false
	}

	client, err := utils.CreateIdentityServicesClient(fmt.Sprintf("%s:%d", domainDefinition.IdentityServiceHost, domainDefinition.IdentityServicePort))
	if err != nil {
		return false
	}

	verifyIdentity, err := client.GetIdentity(context.Background(), ident.Address)
	if err != nil {
		return false
	}

	if verifyIdentity.IdentityPublicKey != ident.IdentityPublicKey {
		return false
	}

	return true

}



func VerifyIdentityAttributes(ident protocol.Identity) IdentityAttributeVerificationResult {

	// Index attestations
	var attestIndex = make(map[string][]*protocol.Identity_Attestation)
	for _, attest := range ident.Attestations {
		attestIndex[attest.AttributeName] = append(attestIndex[attest.AttributeName], attest)
	}

	var result IdentityAttributeVerificationResult

	result.ValidAttestedAttributes = make(map[string]*protocol.Identity_Attribute)
	result.ValidAttestors = make(map[string][]*protocol.Address)
	result.InvalidAttestedAttributes = make(map[string]*protocol.Identity_Attribute)
	result.InvalidAttestors = make(map[string][]*protocol.Address)

	for key, attr := range ident.Attributes {

		// is there an attestation for this attribute?
		if attestIndex[key] != nil {

			for _, att := range attestIndex[key] {

				// Include domain and identity of attestors?
				attestorAddress := protocol.Address{
					Domain:att.AttestorDomain,
					DescriptorReference:att.AttestorDescriptor,
				}

				if VerifyAttestation(attr.ValueHash, att) {
					result.ValidAttestedAttributes[key] = attr
					result.ValidAttestors[key] = append(result.ValidAttestors[key], &attestorAddress)
				} else {
					result.InvalidAttestedAttributes[key] = attr
					result.InvalidAttestors[key] = append(result.InvalidAttestors[key], &attestorAddress)
				}

			}

		}

	}

	return result

}


func VerifyAttestation(encodedHash string, attest *protocol.Identity_Attestation) bool {

	data := sha256.HashData([]byte(attest.AttestorPublicKey + attest.AttributeName + encodedHash))
	decodedPublicKey, err := encoding.Decode(attest.AttestorPublicKey)
	if err != nil {
		return false
	}

	decodedSignature, err := encoding.Decode(attest.Attestation)
	if err != nil {
		return false
	}



	return ed25519.Verify(decodedPublicKey, decodedSignature, data)

}