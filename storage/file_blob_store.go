package storage

import (
	"github.com/datumchi/go/crypto/hsm"
	"github.com/datumchi/go/utility/logger"
	"io/ioutil"
)

type FileBlobStore struct {

	BasePath string
	Crypto   *hsm.SecurityModule

}



func CreateFileBlobStore(basePath string) (FileBlobStore, error) {

	store := FileBlobStore {
		BasePath:basePath,
	}

	return store, nil

}

func CreateSecureFileBlobStore(basePath string, encryptionManager *hsm.SecurityModule) (FileBlobStore, error) {

	store := FileBlobStore {
		BasePath:basePath,
		Crypto: encryptionManager,
	}

	return store, nil

}

func (store FileBlobStore) Get(identifier string) ([]byte, error) {

	blobfile := store.BasePath + "/" + identifier + ".fairxblob"
	content, err := ioutil.ReadFile(blobfile)
	if err != nil {
		return nil, err
	}

	var preparedContent []byte
	if store.Crypto != nil {

		preparedContent, err = (*store.Crypto).DecryptData(content)
		if err != nil {
			return nil, err
		}

	} else {
		preparedContent = content
	}

	return preparedContent, nil

}



func (store FileBlobStore) Save(identifier string, data []byte) error {

	var preparedData []byte
	var err error
	if store.Crypto != nil {

		preparedData, err = (*store.Crypto).EncryptData(data)
		if err != nil {
			return err
		}

	} else {
		preparedData = data
	}

	blobfile := store.BasePath + "/" + identifier + ".fairxblob"
	err = ioutil.WriteFile(blobfile, preparedData, 0644)
	if err != nil {
		logger.Errorf("There was a problem writing out to file:  %v", err)
		return err
	}

	return nil

}