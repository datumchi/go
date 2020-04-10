package storage


type BlobStore interface {

	Get(identifier string) ([]byte, error)
	Save(identifier string, data []byte) error

}
