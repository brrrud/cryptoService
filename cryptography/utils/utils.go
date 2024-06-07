package utils

// rotl32 выполняет циклический сдвиг влево
func Rotl32(x, y uint32) uint32 {
	return (x << (y & 31)) | (x >> (32 - (y & 31)))
}

// rotr32 выполняет циклический сдвиг вправо
func Rotr32(x, y uint32) uint32 {
	return (x >> (y & 31)) | (x << (32 - (y & 31)))
}
