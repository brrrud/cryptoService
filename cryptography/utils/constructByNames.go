package utils

import (
	"cryptoService/cryptography/algorithms"
	"cryptoService/cryptography/modes"
	"cryptoService/cryptography/paddings"
	"errors"
)

func GetAlgorithmByName(nameAlgorithm string, key []byte) algorithms.SymmetricEncryption {
	switch nameAlgorithm {
	case "RC6":
		return algorithms.NewCipher(key)
	case "Camellia":
		ans, _ := algorithms.NewCamellia(key)
		return ans
	case "TwoFish":
		ans, _ := algorithms.NewTwoFishCipher(key)
		return ans
	case "Serpent":
		ans, _ := algorithms.NewSerpentCipher(key)
		return ans
	}
	return nil
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

func GetModeByName(nameMode string, pad paddings.Padding, algo algorithms.SymmetricEncryption) (modes.CipherMode, error) {
	switch nameMode {
	case "CBC":
		mode, err := modes.NewCBC(algo, pad)
		if err != nil {
			return nil, err
		}
		return mode, nil
	case "CFB":
		mode, err := modes.NewCFBMode(algo, pad)
		if err != nil {
			return nil, err
		}
		return mode, nil
	case "CTR":
		mode, err := modes.NewCTRMode(algo)
		if err != nil {
			return nil, err
		}
		return mode, nil
	case "ECB":
		mode, err := modes.NewECB(algo, pad)
		if err != nil {
			return nil, err
		}
		return mode, nil
	case "OFB":
		mode, err := modes.NewOFBMode(algo)
		if err != nil {
			return nil, err
		}
		return mode, nil
	case "PCBC":
		mode, err := modes.NewPCBC(algo, pad)
		if err != nil {
			return nil, err
		}
		return mode, nil
	default:
		return nil, errors.New("unsupported mode")
	}
}

func StringToByteArray(s string) []byte {
	return []byte(s)
}

func BytesToString(byteArray []byte) string {
	return string(byteArray)
}
