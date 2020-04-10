package domain_test

import (
	"github.com/datumchi/go/naming/dns"
	"github.com/datumchi/go/verifier/domain"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Verification", func() {

	Describe("Domain Verification", func() {

		Context("Using a correct domain definition", func() {

			var domainDefinition dns.DomainDefinition


			BeforeEach(func() {

				domainDefinition = dns.DomainDefinition {
					Domain:"alpha.fairx.io",
					PublicKey:"Z5UWZC4ED7IQUUDYCSBU2WGHAXASHUTHQN66X7CZJX7IZL6AEY2Q",
					Signature:"HOWRNFPJD432AQTAAVYNNLNLTRM4TVHCADJCQFTYBRUH5LRDZJI4M2XQ4VTPOSFU2NDVUE3GQOGFVUB7OQWZTSWLEO7GLYRIZ33WYBQ",
					CollabServiceHost:"alphacollab.fairx.io",
					CollabServicePort:17177,
					IdentityServiceHost:"alphaidentity.fairx.io",
					IdentityServicePort:17117,
				}

			})

			It("Should verify the domain definition as valid", func() {

				isValid := domain.VerifyDomainDefinition(domainDefinition)
				Expect(isValid).To(BeTrue())


			})




		})




	})



})
