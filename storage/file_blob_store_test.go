package storage_test

import (
	"github.com/datumchi/go/crypto/hsm"
	"github.com/datumchi/go/storage"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"os"
)

var _ = Describe("FileBlobStore", func() {


	Describe("Getting a blob in a file after saving it encrypted", func() {

		var key *[32]byte
		var fileBlockStore storage.FileBlobStore
		var err error
		var basePath string = "/tmp"
		var testIdentifier string = "TEST"

		BeforeEach(func() {
			key = hsm.NewEncryptionKey()
			simpleModule, _ := hsm.CreateSimpleSecurityModule(key)
			fileBlockStore, err = storage.CreateSecureFileBlobStore(basePath, &simpleModule)
			if err != nil {
				Fail("There was an error creating a secure file blob store: " + err.Error())
			}
		})

		AfterEach(func() {
			err := os.Remove(basePath + "/" + testIdentifier + ".fairxblob")
			if err != nil {
				Fail("Unable to remove file:  " + err.Error())
			}
		})

		It("Encrypted blobs should be retrieved after saved", func() {

			var content []byte
			err = fileBlockStore.Save(testIdentifier, []byte("Encrypted Test Data"))
			Expect(err).To(BeNil())

			content, err = fileBlockStore.Get(testIdentifier)
			Expect(err).To(BeNil())
			Expect(content).ToNot(BeNil())
			Expect(content).To(Equal([]byte("Encrypted Test Data")))

		})

	})

	Describe("Getting a blob in a file after saving it unencrypted", func() {

		var fileBlockStore storage.FileBlobStore
		var err error
		var basePath string = "/tmp"
		var testIdentifier string = "TEST"

		BeforeEach(func() {
			fileBlockStore, err = storage.CreateFileBlobStore(basePath)
			if err != nil {
				Fail("There was an error creating a file blob store: " + err.Error())
			}
		})

		AfterEach(func() {
			err := os.Remove(basePath + "/" + testIdentifier + ".fairxblob")
			if err != nil {
				Fail("Unable to remove file:  " + err.Error())
			}
		})

		It("Blobs should be retrieved aster saved", func() {

			err = fileBlockStore.Save(testIdentifier, []byte("This is a get test"))
			Expect(err).To(BeNil())

			content, err := fileBlockStore.Get(testIdentifier)
			Expect(err).To(BeNil())
			Expect(content).ToNot(BeNil())
			Expect(content).To(Equal([]byte("This is a get test")))

		})

	})


	Describe("Saving a blob in a file unencrypted", func() {

		var fileBlockStore storage.FileBlobStore
		var err error
		var basePath string = "/tmp"
		var testIdentifier string = "TEST"

		BeforeEach(func() {
			fileBlockStore, err = storage.CreateFileBlobStore(basePath)
			if err != nil {
				Fail("There was an error creating a file blob store: " + err.Error())
			}
		})

		AfterEach(func() {
			err := os.Remove(basePath + "/" + testIdentifier + ".fairxblob")
			if err != nil {
				Fail("Unable to remove file:  " + err.Error())
			}
		})

		It("Should save the text as a byte array unencrypted", func() {

			err := fileBlockStore.Save(testIdentifier, []byte("this is a test"))
			Expect(err).To(BeNil())

			content, err := ioutil.ReadFile(basePath + "/" + testIdentifier + ".fairxblob")
			Expect(err).To(BeNil())
			Expect(content).To(BeEquivalentTo([]byte("this is a test")))

		})

	})


	Describe("Saving a blob in a file encrypted", func() {

		var key *[32]byte
		var fileBlockStore storage.FileBlobStore
		var err error
		var basePath string = "/tmp"
		var testIdentifier string = "TEST"

		BeforeEach(func() {
			key = hsm.NewEncryptionKey()
			simpleModule, _ := hsm.CreateSimpleSecurityModule(key)
			fileBlockStore, err = storage.CreateSecureFileBlobStore(basePath, &simpleModule)
			if err != nil {
				Fail("There was an error creating a secure file blob store: " + err.Error())
			}
		})

		AfterEach(func() {
			err := os.Remove(basePath + "/" + testIdentifier + ".fairxblob")
			if err != nil {
				Fail("Unable to remove file:  " + err.Error())
			}
		})

		It("Should save the text as a byte array encrypted", func() {

			err := fileBlockStore.Save(testIdentifier, []byte("This is an encrypted test"))
			Expect(err).To(BeNil())

			encryptedData, err := ioutil.ReadFile(basePath + "/" + testIdentifier + ".fairxblob")
			Expect(err).To(BeNil())
			Expect(encryptedData).ToNot(BeNil())

			testDecryptedData, err := hsm.Decrypt(encryptedData, key)
			Expect(err).To(BeNil())
			Expect(testDecryptedData).ToNot(BeNil())
			Expect(testDecryptedData).To(Equal([]byte("This is an encrypted test")))

		})

	})


})
