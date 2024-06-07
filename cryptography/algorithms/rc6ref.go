package algorithms

import (
	"encoding/binary"
	"errors"
)

const blockSize = 16

func (r Rc6Cipher) GetBlockSize() int {
	return blockSize
}

var defaultRounds uint32 = 20

var constP uint32 = 0xB7E15163
var constQ uint32 = 0x9E3779B9

type Rc6Cipher struct {
	rounds   uint32
	keyLen   uint8
	key      []uint32
	keySched []uint32
}

func NewCipher(key []byte) Rc6Cipher {
	var c Rc6Cipher
	for len(key)%4 != 0 {
		key = append(key, 0)
	}
	c.rounds = defaultRounds
	c.key = make([]uint32, len(key)/4)
	var a, b, j, i, s, v, klen, keySchedMax uint32
	klen = uint32(len(key))
	for i = 0; i < klen; i += 4 {
		c.key[i/4] = binary.LittleEndian.Uint32(key[i:])
	}
	keySchedMax = 2*c.rounds + 4
	c.keyLen = uint8(len(c.key))
	c.keySched = make([]uint32, keySchedMax)
	c.keySched[0] = constP
	for i = 1; i < keySchedMax; i++ {
		c.keySched[i] = c.keySched[i-1] + constQ
	}
	v = 3 * (2*c.rounds + 4)
	if uint32(c.keyLen) > (2*c.rounds + 4) {
		v = uint32(c.keyLen)
	}
	a = 0
	b = 0
	i = 0
	j = 0
	for s = 1; s <= v; s++ {
		c.keySched[i] = rotl32(c.keySched[i]+a+b, 3)
		a = c.keySched[i]
		c.key[j] = rotl32(c.key[j]+a+b, a+b)
		b = c.key[j]
		i = (i + 1) % (2*c.rounds + 4)
		j = (j + 1) % uint32(c.keyLen)
	}
	return c
}

func (Rc6Cipher) BlockSize() int {
	return int(blockSize)
}

func (r Rc6Cipher) GetKeySched() []uint32 {
	return r.keySched
}

func (r Rc6Cipher) EncryptBlock(dst, src []byte) {
	if len(src) != blockSize {
		panic("Incorrect amount of data passed to Encrypt")
	}
	ct := make([]uint32, 4)
	for i := 0; i < blockSize; i += 4 {
		ct[i/4] = binary.LittleEndian.Uint32(src[i:])
	}
	var a, b, c, d, t, u, i, x uint32
	a = ct[0]
	b = ct[1]
	c = ct[2]
	d = ct[3]
	b = b + r.keySched[0]
	d = d + r.keySched[1]
	for i = 1; i <= r.rounds; i++ {

		t = rotl32((b * (2*b + 1)), 5)
		u = rotl32((d * (2*d + 1)), 5)
		a = rotl32((a^t), u) + r.keySched[2*i]
		c = rotl32((c^u), t) + r.keySched[2*i+1]
		x = a
		a = b
		b = c
		c = d
		d = x
	}
	a = a + r.keySched[2*r.rounds+2]
	c = c + r.keySched[2*r.rounds+3]
	binary.LittleEndian.PutUint32(dst[0:], a)
	binary.LittleEndian.PutUint32(dst[4:], b)
	binary.LittleEndian.PutUint32(dst[8:], c)
	binary.LittleEndian.PutUint32(dst[12:], d)

}

func (r Rc6Cipher) DecryptBlock(dst, src []byte) {
	if len(src) != blockSize {
		panic("Incorrect amount of data passed to Encrypt")
	}
	ct := make([]uint32, 4)
	for i := 0; i < blockSize; i += 4 {
		ct[i/4] = binary.LittleEndian.Uint32(src[i:])
	}
	var a, b, c, d, t, u, i, x uint32
	a = ct[0]
	b = ct[1]
	c = ct[2]
	d = ct[3]
	c = c - r.keySched[2*r.rounds+3]
	a = a - r.keySched[2*r.rounds+2]
	for i = r.rounds; i >= 1; i-- {
		x = d
		d = c
		c = b
		b = a
		a = x
		u = rotl32((d * (2*d + 1)), 5)
		t = rotl32((b * (2*b + 1)), 5)
		c = rotr32(c-r.keySched[2*i+1], t) ^ u
		a = rotr32(a-r.keySched[2*i], u) ^ t
	}
	d = d - r.keySched[1]
	b = b - r.keySched[0]
	binary.LittleEndian.PutUint32(dst[0:], a)
	binary.LittleEndian.PutUint32(dst[4:], b)
	binary.LittleEndian.PutUint32(dst[8:], c)
	binary.LittleEndian.PutUint32(dst[12:], d)
}

func rotl32(x, y uint32) uint32 {
	var w uint32 = 32
	return ((x) << (y & (w - 1))) | ((x) >> (w - (y & (w - 1))))
}

func rotr32(x, y uint32) uint32 {
	var w uint32 = 32
	return ((x) >> (y & (w - 1))) | ((x) << (w - (y & (w - 1))))
}
func (r Rc6Cipher) Encrypt(plaintext []byte) ([]byte, error) {
	if len(plaintext)%blockSize != 0 {
		return nil, errors.New("plaintext length must be a multiple of block size")
	}
	cipherText := make([]byte, len(plaintext))
	for i := 0; i < len(plaintext); i += blockSize {
		r.EncryptBlock(cipherText[i:i+blockSize], plaintext[i:i+blockSize])
	}
	return cipherText, nil
}

func (r Rc6Cipher) Decrypt(cipherText []byte) ([]byte, error) {
	if len(cipherText)%blockSize != 0 {
		return nil, errors.New("cipherText length must be a multiple of block size")
	}
	plainText := make([]byte, len(cipherText))
	for i := 0; i < len(cipherText); i += blockSize {
		r.DecryptBlock(plainText[i:i+blockSize], cipherText[i:i+blockSize])
	}
	return plainText, nil
}
