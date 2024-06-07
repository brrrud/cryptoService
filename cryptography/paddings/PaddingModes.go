package paddings

type PaddingMode int

const (
	Zeros PaddingMode = iota
	ANSIX923
	PKCS7
	ISO10126
)
