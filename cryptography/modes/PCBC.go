package modes

import (
	"cryptoService/cryptography/algorithms"
	"cryptoService/cryptography/paddings"
	"errors"
)

type PCBC struct {
	cipher    algorithms.SymmetricEncryption
	padding   paddings.Padding
	blockSize int
}

func NewPCBC(cipher algorithms.SymmetricEncryption, padding paddings.Padding) (*PCBC, error) {
	blockSize := cipher.GetBlockSize()
	if blockSize <= 0 {
		return nil, errors.New("invalid block size")
	}
	return &PCBC{
		cipher:    cipher,
		padding:   padding,
		blockSize: blockSize,
	}, nil
}
func (pcbc *PCBC) Encrypt(plaintext []byte) ([]byte, error) {
	padded, err := pcbc.padding.Pad(plaintext, pcbc.blockSize)
	if err != nil {
		return nil, err
	}

	iv, err := GenerateIV(pcbc.blockSize)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, 0, len(padded)+len(iv))
	ciphertext = append(ciphertext, iv...)

	prevCipherBlock := iv
	for i := 0; i < len(padded); i += pcbc.blockSize {
		block := padded[i : i+pcbc.blockSize]
		xoredBlock := xorBytes(block, prevCipherBlock)

		encryptedBlock, err := pcbc.cipher.Encrypt(xoredBlock)
		if err != nil {
			return nil, err
		}

		ciphertext = append(ciphertext, encryptedBlock...)
		prevCipherBlock = xorBytes(block, encryptedBlock)
	}

	return ciphertext, nil
}

func (pcbc *PCBC) Decrypt(ciphertext []byte) ([]byte, error) {
	if len(ciphertext) < pcbc.blockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := ciphertext[:pcbc.blockSize]
	ciphertext = ciphertext[pcbc.blockSize:]

	plaintext := make([]byte, 0, len(ciphertext))
	prevCipherBlock := iv
	for i := 0; i < len(ciphertext); i += pcbc.blockSize {
		block := ciphertext[i : i+pcbc.blockSize]

		decryptedBlock, err := pcbc.cipher.Decrypt(block)
		if err != nil {
			return nil, err
		}

		xoredBlock := xorBytes(decryptedBlock, prevCipherBlock)
		plaintext = append(plaintext, xoredBlock...)

		prevCipherBlock = xorBytes(xoredBlock, block)
	}

	return pcbc.padding.Unpad(plaintext, pcbc.blockSize)
}
