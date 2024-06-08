package paddings

import "fmt"

type ANSIX923Padding struct{}

func NewANSIX923Padding() *ANSIX923Padding {
	return &ANSIX923Padding{}
}

func (ap ANSIX923Padding) Pad(data []byte, blockSize int) ([]byte, error) {
	paddingSize := blockSize - len(data)%blockSize
	padding := make([]byte, paddingSize)
	padding[paddingSize-1] = byte(paddingSize)
	paddedData := append(data, padding...)
	return paddedData, nil
}
func (ap ANSIX923Padding) Unpad(data []byte, blockSize int) ([]byte, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("invalid padding size")
	}
	paddingSize := int(data[len(data)-1])
	if paddingSize <= 0 || paddingSize > len(data) || paddingSize > blockSize {
		return nil, fmt.Errorf("invalid padding size")
	}
	for i := 0; i < paddingSize-1; i++ {
		if data[len(data)-2-i] != 0 { // ИСПРАВЛЕНО: data[len(data)-1-i] -> data[len(data)-2-i]
			return nil, fmt.Errorf("invalid padding")
		}
	}
	return data[:len(data)-paddingSize], nil
}
