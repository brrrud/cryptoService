package algorithms

type SymmetricEncryption interface {
	Encrypt(plaintext []byte) ([]byte, error)
	Decrypt(cipherText []byte) ([]byte, error)
	GetBlockSize() int
}
