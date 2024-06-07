package paddings

import (
	"crypto/rand"
	"fmt"
)

type ISO10126Padding struct{}

func NewISO10126Padding() *ISO10126Padding {
	return &ISO10126Padding{}
}

func (ip ISO10126Padding) Pad(data []byte, blockSize int) ([]byte, error) {
	paddingSize := blockSize - len(data)%blockSize
	padding := make([]byte, paddingSize-1)
	_, err := rand.Read(padding)
	if err != nil {
		return nil, err
	}
	padding = append(padding, byte(paddingSize))
	paddedData := append(data, padding...)
	return paddedData, nil
}

func (ip ISO10126Padding) Unpad(data []byte, blockSize int) ([]byte, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("invalid padding size")
	}
	paddingSize := int(data[len(data)-1])
	if paddingSize > len(data) || paddingSize > blockSize {
		return nil, fmt.Errorf("invalid padding size")
	}
	return data[:len(data)-paddingSize], nil
}
