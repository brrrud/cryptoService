package paddings

import (
	"errors"
	"fmt"
)

type ZerosPadding struct{}

func NewZerosPadding() *ZerosPadding {
	return &ZerosPadding{}
}

// Pad - метод для добавления нулевой набивки
func (zp ZerosPadding) Pad(data []byte, blockSize int) ([]byte, error) {
	paddingSize := blockSize - len(data)%blockSize
	paddedData := append(data, make([]byte, paddingSize)...)
	return paddedData, nil
}

// Unpad - метод для удаления нулевой набивки
func (zp ZerosPadding) Unpad(data []byte, blockSize int) ([]byte, error) {
	if len(data)%blockSize != 0 {
		return nil, fmt.Errorf("invalid padding: data length is not a multiple of block size")
	}

	for i := len(data) - 1; i >= 0; i-- {
		if data[i] != 0 {
			return data[:i+1], nil
		}
	}

	return nil, errors.New("invalid padding: all bytes are zero")
}
