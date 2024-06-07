package modes

import (
	"cryptoService/cryptography/algorithms"
	"cryptoService/cryptography/paddings"
)

type CBCMode struct {
	blockCipher algorithms.SymmetricEncryption
	padding     paddings.Padding
	iv          []byte
}

func NewCBC(blockCipher algorithms.SymmetricEncryption, padding paddings.Padding) (*CBCMode, error) {
	blockSize := blockCipher.GetBlockSize()
	iv, err := GenerateIV(blockSize)
	if err != nil {
		return nil, err
	}
	return &CBCMode{blockCipher: blockCipher, padding: padding, iv: iv}, nil
}

func (c *CBCMode) Encrypt(plaintext []byte) ([]byte, error) {
	blockSize := c.blockCipher.GetBlockSize()
	paddedPlaintext, _ := c.padding.Pad(plaintext, blockSize)

	previousBlock := c.iv
	ciphertext := make([]byte, 0, len(paddedPlaintext))

	for i := 0; i < len(paddedPlaintext); i += blockSize {
		block := paddedPlaintext[i : i+blockSize]
		for j := 0; j < blockSize; j++ {
			block[j] ^= previousBlock[j]
		}
		encryptedBlock, _ := c.blockCipher.Encrypt(block)
		ciphertext = append(ciphertext, encryptedBlock...)
		previousBlock = encryptedBlock
	}

	return ciphertext, nil
}

func (c *CBCMode) Decrypt(ciphertext []byte) ([]byte, error) {
	blockSize := c.blockCipher.GetBlockSize()

	previousBlock := c.iv
	plaintext := make([]byte, 0, len(ciphertext))

	for i := 0; i < len(ciphertext); i += blockSize {
		block := ciphertext[i : i+blockSize]
		decryptedBlock, _ := c.blockCipher.Decrypt(block)
		for j := 0; j < blockSize; j++ {
			decryptedBlock[j] ^= previousBlock[j]
		}
		plaintext = append(plaintext, decryptedBlock...)
		previousBlock = block
	}

	unpaddedPlaintext, _ := c.padding.Unpad(plaintext, blockSize)
	return unpaddedPlaintext, nil
}

//type CBC struct {
//	blockCipher algorithms.SymmetricEncryption
//	padding     paddings.Padding
//	blockSize   int
//}
//
//// NewCBC создает новый экземпляр CBC
//func NewCBC(blockCipher algorithms.SymmetricEncryption, padding paddings.Padding) *CBC {
//	return &CBC{
//		blockCipher: blockCipher,
//		padding:     padding,
//		blockSize:   blockCipher.GetBlockSize(),
//	}
//}
//
//// Encrypt шифрует данные в режиме CBC
//// Encrypt шифрует данные в режиме CBC
//func (c *CBC) Encrypt(plaintext []byte) ([]byte, error) {
//	paddedData, err := c.padding.Pad(plaintext, c.blockSize)
//	if err != nil {
//		return nil, err
//	}
//
//	cipherText := make([]byte, c.blockSize+len(paddedData))
//	iv := cipherText[:c.blockSize]
//
//	// Генерация псевдослучайного IV
//	seed := time.Now().UnixNano()
//	for i := range iv {
//		seed = (seed*9301 + 49297) % 233280
//		iv[i] = byte(seed % 256)
//	}
//
//	previousBlock := iv
//	for i := 0; i < len(paddedData); i += c.blockSize {
//		block := paddedData[i : i+c.blockSize]
//		for j := 0; j < c.blockSize; j++ {
//			block[j] ^= previousBlock[j]
//		}
//		encryptedBlock, err := c.blockCipher.Encrypt(block)
//		if err != nil {
//			return nil, err
//		}
//		copy(cipherText[c.blockSize+i:], encryptedBlock)
//		previousBlock = encryptedBlock
//	}
//
//	return cipherText, nil
//}
//
//func (c *CBC) Decrypt(cipherText []byte) ([]byte, error) {
//	if len(cipherText) < c.blockSize {
//		return nil, errors.New("ciphertext too short")
//	}
//
//	iv := cipherText[:c.blockSize]
//	cipherText = cipherText[c.blockSize:]
//
//	if len(cipherText)%c.blockSize != 0 {
//		return nil, errors.New("ciphertext is not a multiple of the block size")
//	}
//
//	previousBlock := iv
//	decryptedText := make([]byte, len(cipherText))
//
//	for i := 0; i < len(cipherText); i += c.blockSize {
//		block := cipherText[i : i+c.blockSize]
//		decryptedBlock, err := c.blockCipher.Decrypt(block)
//		if err != nil {
//			return nil, err
//		}
//		for j := 0; j < c.blockSize; j++ {
//			decryptedBlock[j] ^= previousBlock[j]
//		}
//		copy(decryptedText[i:], decryptedBlock)
//		previousBlock = block
//	}
//
//	return c.padding.Unpad(decryptedText, c.blockSize)
//}
