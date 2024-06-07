package modes

import (
	"cryptoService/cryptography/algorithms"
	"cryptoService/cryptography/paddings"
	"errors"
)

type CFBMode struct {
	blockCipher algorithms.SymmetricEncryption
	padding     paddings.Padding
	iv          []byte
}

func NewCFBMode(blockCipher algorithms.SymmetricEncryption, padding paddings.Padding) (*CFBMode, error) {
	blockSize := blockCipher.GetBlockSize()
	iv, err := GenerateIV(blockSize)
	if err != nil {
		return nil, err
	}

	return &CFBMode{
		blockCipher: blockCipher,
		padding:     padding,
		iv:          iv,
	}, nil
}

// Encrypt encrypts the plaintext using CFB mode.
func (cfb *CFBMode) Encrypt(plaintext []byte) ([]byte, error) {
	blockSize := cfb.blockCipher.GetBlockSize()
	paddedPlaintext, err := cfb.padding.Pad(plaintext, blockSize)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, len(cfb.iv)+len(paddedPlaintext))
	copy(ciphertext[:blockSize], cfb.iv)

	prevCipherBlock := cfb.iv
	for i := 0; i < len(paddedPlaintext); i += blockSize {
		encryptedBlock, err := cfb.blockCipher.Encrypt(prevCipherBlock)
		if err != nil {
			return nil, err
		}

		block := xorBytes(paddedPlaintext[i:i+blockSize], encryptedBlock)
		copy(ciphertext[blockSize+i:], block)

		prevCipherBlock = block
	}

	return ciphertext, nil
}

// Decrypt decrypts the ciphertext using CFB mode.
func (cfb *CFBMode) Decrypt(ciphertext []byte) ([]byte, error) {
	blockSize := cfb.blockCipher.GetBlockSize()
	if len(ciphertext) < blockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := ciphertext[:blockSize]
	ciphertext = ciphertext[blockSize:]

	plaintext := make([]byte, len(ciphertext))
	prevCipherBlock := iv
	for i := 0; i < len(ciphertext); i += blockSize {
		encryptedBlock, err := cfb.blockCipher.Encrypt(prevCipherBlock)
		if err != nil {
			return nil, err
		}

		block := xorBytes(ciphertext[i:i+blockSize], encryptedBlock)
		copy(plaintext[i:], block)

		prevCipherBlock = ciphertext[i : i+blockSize]
	}

	unpaddedPlaintext, err := cfb.padding.Unpad(plaintext, blockSize)
	if err != nil {
		return nil, err
	}

	return unpaddedPlaintext, nil
}
