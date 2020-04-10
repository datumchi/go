package identity_test

import (
	"crypto/rand"
	"github.com/datumchi/go/crypto/sha256"
	"github.com/datumchi/go/encoding"
	"github.com/datumchi/go/generated/protocol"
	"github.com/datumchi/go/verifier/identity"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/crypto/ed25519"
)

var _ = Describe("Verification", func() {


	Describe("Identity Attribute Attestation Verification", func() {

		Context("A valid identity with valid attestations of attributes", func() {

			var validIdentity protocol.Identity

			BeforeEach(func() {

				var identityAttestationList []*protocol.Identity_Attestation
				publicKey, privateKey, _ := ed25519.GenerateKey(rand.Reader)
				encodedPublicKey := encoding.Encode(publicKey)

				address := protocol.Address{
					Domain: "alpha.fairx.io",
					DescriptorReference:"test@fairx.io",
				}

				givenNameHash := encoding.Encode(sha256.HashData([]byte("Test")))
				givenNameAttr := protocol.Identity_Attribute{
					Name:        "given_name",
					ValueHash:   givenNameHash,
					Description: "Given Name",
				}
				givenNameAttestationSignature := ed25519.Sign(privateKey, sha256.HashData([]byte(encodedPublicKey + "given_name" + givenNameHash)))
				givenNameAttestation := protocol.Identity_Attestation{
					AttributeName:      "given_name",
					AttestorDomain:     address.Domain,
					AttestorDescriptor: address.DescriptorReference,
					AttestorPublicKey:  encodedPublicKey,
					Attestation:        encoding.Encode(givenNameAttestationSignature),
				}

				surNameHash := encoding.Encode(sha256.HashData([]byte("User")))
				surNameAttr := protocol.Identity_Attribute{
					Name:        "sur_name",
					ValueHash:   surNameHash,
					Description: "Surname/Last Name/Family Name",
				}
				surNameAttestationSignature := ed25519.Sign(privateKey, sha256.HashData([]byte(encodedPublicKey + "sur_name" + surNameHash)))
				surNameAttestation := protocol.Identity_Attestation{
					AttributeName:      "sur_name",
					AttestorDomain:     address.Domain,
					AttestorDescriptor: address.DescriptorReference,
					AttestorPublicKey:  encodedPublicKey,
					Attestation:        encoding.Encode(surNameAttestationSignature),
				}


				attributeMap := map[string]*protocol.Identity_Attribute {
					"given_name": &givenNameAttr,
					"sur_name": &surNameAttr,
				}

				identityAttestationList = append(identityAttestationList, &givenNameAttestation)
				identityAttestationList = append(identityAttestationList, &surNameAttestation)

				validIdentity = protocol.Identity{
					Address:           &address,
					IdentityPublicKey: encodedPublicKey,
					Attributes:        attributeMap,
					Attestations:      identityAttestationList,
				}

			})

			It("Returns a full list of verified attributes with no invalid attributes", func() {

				//result := identity.VerifyIdentityAttributes(validIdentity)
				//Expect(len(result.ValidAttestedAttributes)).To(Equal(2))



			})





		})





	})



})
