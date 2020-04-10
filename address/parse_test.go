package address_test

import (
	"github.com/datumchi/go/address"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parse", func() {

	Describe("Parse FairX Addresses", func() {

		Context("Correctly Formatted Address with Path Descriptor", func() {

			var validAddress = "[test@fairx.io*/path/to/something]alpha.fairx.io"

			It("Should parse correctly", func() {

				addr, err := address.ToAddress(validAddress)
				Expect(err).To(BeNil())
				Expect(addr.Domain).To(Equal("alpha.fairx.io"))
				Expect(addr.DescriptorReference).To(Equal("test@fairx.io"))
				Expect(addr.DescriptorPath).To(Equal("/path/to/something"))

			})


		})

		Context("Correctly Formatted Address without Path Descriptor", func() {

			var validAddress = "[test@fairx.io]alpha.fairx.io"

			It("Should parse correctly", func() {

				addr, err := address.ToAddress(validAddress)
				Expect(err).To(BeNil())
				Expect(addr.Domain).To(Equal("alpha.fairx.io"))
				Expect(addr.DescriptorReference).To(Equal("test@fairx.io"))
				Expect(addr.DescriptorPath).To(BeEmpty())

			})


		})

		Context("Incorrectly Formatted Address", func() {

			var validAddress = "test@fairx.io]alpha.fairx.io"

			It("Should not parse and return an error", func() {

				_, err := address.ToAddress(validAddress)
				Expect(err).ToNot(BeNil())

			})


		})



	})



})
