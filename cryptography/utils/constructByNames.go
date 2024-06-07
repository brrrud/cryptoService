package utils

import (
	"algorithms/cryptography/algorithms"
	"algorithms/cryptography/modes"
	"algorithms/cryptography/paddings"
	"errors"
	"fmt"
)

func GetAlgorithmByName(nameAlgorithm string) algorithms.SymmetricEncryption {
	switch nameAlgorithm {
	case "RC6":
		return &algorithms.Rc6Cipher{}
	case "Camellia":
		fmt.Println("Not available")
	}
	return nil
}

func GetModeByName(nameMode string) modes.CipherMode {
	switch nameMode {
	case "CBC":
		return &modes.CBCMode{}
	case "CFB":
		return &modes.CFBMode{}
	case "CTR":
		return &modes.CTRMode{}
	case "ECB":
		return &modes.ECB{}
	case "OFB":
		return &modes.OFBMode{}
	case "PCBC":
		return &modes.PCBC{}
	default:
		return nil
	}
}

func GetPaddingByName(namePadding string) paddings.Padding {
	switch namePadding {
	case "ANSIX923":
		return &paddings.ANSIX923Padding{}
	case "ISO10126":
		return &paddings.ISO10126Padding{}
	case "PKCS7":
		return &paddings.PKCS7Padding{}
	case "ZEROS":
		return paddings.ZerosPadding{}
	default:
		return nil
	}
}

type CipherModeBuilder struct {
	encryptionAlgorithm algorithms.SymmetricEncryption
	mode                modes.CipherMode
	padding             paddings.Padding
	key                 string
}

func NewCipherModeBuilder() *CipherModeBuilder {
	return &CipherModeBuilder{}
}

func (b *CipherModeBuilder) SetEncryptionAlgorithm(algorithm algorithms.SymmetricEncryption) *CipherModeBuilder {
	b.encryptionAlgorithm = algorithm
	return b
}

func (b *CipherModeBuilder) SetMode(mode modes.CipherMode) *CipherModeBuilder {
	b.mode = mode
	return b
}

func (b *CipherModeBuilder) SetPadding(padding paddings.Padding) *CipherModeBuilder {
	b.padding = padding
	return b
}

func (b *CipherModeBuilder) SetKey(key string) *CipherModeBuilder {
	b.key = key
	return b
}

func requirePaddings(modeName modes.CipherMode) bool {
	switch modeName {
	case &modes.CBCMode{}, &modes.CFBMode{}, &modes.ECB{}, &modes.PCBC{}:
		return true
	default:
		return false
	}
}

func (b *CipherModeBuilder) Build() (modes.CipherMode, error) {
	if b.encryptionAlgorithm == nil {
		return nil, errors.New("encryption algorithm is not set")
	}
	if b.mode == nil {
		return nil, errors.New("mode is not set")
	}
	if requirePaddings(b.mode) && b.padding == nil {
		return nil, errors.New("padding is required but not set")
	}
	return b.mode, nil
}
