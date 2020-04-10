package hsm


type SecurityModule interface {

	EncryptData(data []byte) ([]byte, error)
	DecryptData(data []byte) ([]byte, error)

}
