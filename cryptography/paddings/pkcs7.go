package paddings

import "fmt"

type PKCS7Padding struct{}

func NewPKCS7Padding() *PKCS7Padding {
	return &PKCS7Padding{}
}

func (pp PKCS7Padding) Pad(data []byte, blockSize int) ([]byte, error) {
	paddingSize := blockSize - len(data)%blockSize
	padding := make([]byte, paddingSize)
	for i := range padding {
		padding[i] = byte(paddingSize)
	}
	paddedData := append(data, padding...)
	return paddedData, nil
}

func (pp PKCS7Padding) Unpad(data []byte, blockSize int) ([]byte, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("invalid padding size")
	}
	paddingSize := int(data[len(data)-1])
	if paddingSize > len(data) || paddingSize > blockSize {
		return nil, fmt.Errorf("invalid padding size")
	}
	for i := 0; i < paddingSize; i++ {
		if data[len(data)-1-i] != byte(paddingSize) {
			return nil, fmt.Errorf("invalid padding")
		}
	}
	return data[:len(data)-paddingSize], nil
}
