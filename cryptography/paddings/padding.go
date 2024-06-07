package paddings

type Padding interface {
	Pad(data []byte, blockSize int) ([]byte, error)
	Unpad(data []byte, blockSize int) ([]byte, error)
}
