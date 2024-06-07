package modes

import (
	"cryptoService/cryptography/algorithms"
	"cryptoService/cryptography/paddings"
	"errors"
)

type ECB struct {
	blockCipher algorithms.SymmetricEncryption
	padding     paddings.Padding
	blockSize   int
}

// NewECB создает новый экземпляр ECB
func NewECB(blockCipher algorithms.SymmetricEncryption, padding paddings.Padding) (*ECB, error) {
	return &ECB{
		blockCipher: blockCipher,
		padding:     padding,
		blockSize:   blockCipher.GetBlockSize(),
	}, nil
}

// Encrypt шифрует данные в режиме ECB
func (e *ECB) Encrypt(plaintext []byte) ([]byte, error) {
	paddedData, err := e.padding.Pad(plaintext, e.blockSize)
	if err != nil {
		return nil, err
	}

	var cipherText []byte
	for i := 0; i < len(paddedData); i += e.blockSize {
		block := paddedData[i : i+e.blockSize]
		encryptedBlock, err := e.blockCipher.Encrypt(block)
		if err != nil {
			return nil, err
		}
		cipherText = append(cipherText, encryptedBlock...)
	}

	return cipherText, nil
}

// Decrypt расшифровывает данные в режиме ECB
func (e *ECB) Decrypt(cipherText []byte) ([]byte, error) {
	if len(cipherText)%e.blockSize != 0 {
		return nil, errors.New("invalid ciphertext length")
	}

	var plaintext []byte
	for i := 0; i < len(cipherText); i += e.blockSize {
		block := cipherText[i : i+e.blockSize]
		decryptedBlock, err := e.blockCipher.Decrypt(block)
		if err != nil {
			return nil, err
		}
		plaintext = append(plaintext, decryptedBlock...)
	}

	unpaddedData, err := e.padding.Unpad(plaintext, e.blockSize)
	if err != nil {
		return nil, err
	}

	return unpaddedData, nil
}
