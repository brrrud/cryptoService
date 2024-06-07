package modes

import (
	"cryptoService/cryptography/algorithms"
)

// CTRMode represents the CTR mode of operation.
type CTRMode struct {
	blockCipher algorithms.SymmetricEncryption
	iv          []byte
}

// NewCTRMode creates a new CTR mode instance.
func NewCTRMode(blockCipher algorithms.SymmetricEncryption) (*CTRMode, error) {
	blockSize := blockCipher.GetBlockSize()
	iv, err := GenerateIV(blockSize)
	if err != nil {
		return nil, err
	}
	return &CTRMode{
		blockCipher: blockCipher,
		iv:          iv,
	}, nil
}

// Encrypt encrypts the plaintext using CTR mode.
func (ctr *CTRMode) Encrypt(plaintext []byte) ([]byte, error) {
	blockSize := ctr.blockCipher.GetBlockSize()
	counter := make([]byte, blockSize)
	copy(counter, ctr.iv)

	ciphertext := make([]byte, len(plaintext))
	for i := 0; i < len(plaintext); i += blockSize {
		block := plaintext[i:min(i+blockSize, len(plaintext))]
		streamBlock := ctr.encryptCounter(counter)
		for j := 0; j < len(block); j++ {
			ciphertext[i+j] = block[j] ^ streamBlock[j]
		}
		incrementCounter(counter)
	}
	return ciphertext, nil
}

// Decrypt decrypts the ciphertext using CTR mode.
func (ctr *CTRMode) Decrypt(ciphertext []byte) ([]byte, error) {
	// Decryption in CTR mode is the same as encryption.
	resEncrypt, err := ctr.Encrypt(ciphertext)
	if err != nil {
		return nil, err
	} else {
		return resEncrypt, nil
	}
	//return ctr.Encrypt(ciphertext), nil
}

// encryptCounter encrypts the counter using the block cipher.
func (ctr *CTRMode) encryptCounter(counter []byte) []byte {
	encryptedCounter, err := ctr.blockCipher.Encrypt(counter)
	if err != nil {
		panic(err) // Handle error appropriately in real code.
	}
	return encryptedCounter
}

// incrementCounter increments the counter.
func incrementCounter(counter []byte) {
	for i := len(counter) - 1; i >= 0; i-- {
		counter[i]++
		if counter[i] != 0 {
			break
		}
	}
}
