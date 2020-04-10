package hsm


type SimpleSecurityModule struct {

	Key *[32]byte

}


func CreateSimpleSecurityModule(key *[32]byte) (SecurityModule, error) {

	simpleModule := SimpleSecurityModule{
		Key:key,
	}

	return simpleModule, nil

}


func (mod SimpleSecurityModule) EncryptData(data []byte) ([]byte, error) {
	return Encrypt(data, mod.Key)
}

func (mod SimpleSecurityModule) DecryptData(data []byte) ([]byte, error) {
	return Decrypt(data, mod.Key)
}