package modes

import (
	"cryptoService/cryptography/algorithms"
	"cryptoService/cryptography/paddings"
	"errors"
)

type RandomDelta struct {
	cipher  algorithms.SymmetricEncryption
	padding paddings.Padding
}

func NewRandomDelta(cipher algorithms.SymmetricEncryption, padding paddings.Padding) (*RandomDelta, error) {
	return &RandomDelta{
		cipher:  cipher,
		padding: padding,
	}, nil
}

func (rd *RandomDelta) Encrypt(plaintext []byte) ([]byte, error) {
	blockSize := rd.cipher.GetBlockSize()
	paddedText, err := rd.padding.Pad(plaintext, blockSize)
	if err != nil {
		return nil, err
	}

	var ciphertext []byte
	for i := 0; i < len(paddedText); i += blockSize {
		block := paddedText[i : i+blockSize]
		iv, err := GenerateIV(blockSize)
		if err != nil {
			return nil, err
		}
		xoredBlock := xorBytes(block, iv)
		encryptedBlock, err := rd.cipher.Encrypt(xoredBlock)
		if err != nil {
			return nil, err
		}
		ciphertext = append(ciphertext, iv...)
		ciphertext = append(ciphertext, encryptedBlock...)
	}
	return ciphertext, nil
}

func (rd *RandomDelta) Decrypt(ciphertext []byte) ([]byte, error) {
	blockSize := rd.cipher.GetBlockSize()
	if len(ciphertext)%(2*blockSize) != 0 {
		return nil, errors.New("invalid ciphertext length")
	}

	var plaintext []byte
	for i := 0; i < len(ciphertext); i += 2 * blockSize {
		iv := ciphertext[i : i+blockSize]
		encryptedBlock := ciphertext[i+blockSize : i+2*blockSize]
		decryptedBlock, err := rd.cipher.Decrypt(encryptedBlock)
		if err != nil {
			return nil, err
		}
		originalBlock := xorBytes(decryptedBlock, iv)
		plaintext = append(plaintext, originalBlock...)
	}
	unpaddedText, err := rd.padding.Unpad(plaintext, blockSize)
	if err != nil {
		return nil, err
	}
	return unpaddedText, nil
}
