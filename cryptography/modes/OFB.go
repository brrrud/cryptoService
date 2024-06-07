package modes

import (
	"cryptoService/cryptography/algorithms"
)

// OFBMode is the implementation of the OFB mode.
type OFBMode struct {
	blockCipher algorithms.SymmetricEncryption
	iv          []byte
}

// NewOFBMode creates a new OFB mode with the given block cipher and IV.
func NewOFBMode(blockCipher algorithms.SymmetricEncryption) (*OFBMode, error) {
	iv, err := GenerateIV(blockCipher.GetBlockSize())
	if err != nil {
		return nil, err
	}
	return &OFBMode{
		blockCipher: blockCipher,
		iv:          iv,
	}, nil
}

// Encrypt encrypts plaintext using OFB mode.
func (ofb *OFBMode) Encrypt(plaintext []byte) ([]byte, error) {
	blockSize := ofb.blockCipher.GetBlockSize()
	iv := make([]byte, blockSize)
	copy(iv, ofb.iv)

	ciphertext := make([]byte, len(plaintext))
	streamBlock := make([]byte, blockSize)

	for i := 0; i < len(plaintext); i += blockSize {
		ofb.blockCipher.Encrypt(iv)
		copy(streamBlock, iv)
		for j := 0; j < blockSize && i+j < len(plaintext); j++ {
			ciphertext[i+j] = plaintext[i+j] ^ streamBlock[j]
		}
		copy(iv, streamBlock)
	}

	return ciphertext, nil
}

// Decrypt decrypts ciphertext using OFB mode.
func (ofb *OFBMode) Decrypt(ciphertext []byte) ([]byte, error) {
	// OFB mode encryption and decryption are identical
	return ofb.Encrypt(ciphertext)
}
