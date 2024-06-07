package main

import (
	"cryptoService/cryptography/algorithms"
	"cryptoService/cryptography/modes"
	"cryptoService/cryptography/paddings"
	"fmt"
)

func main() {
	key := []byte("examplekey123456")
	serpent, _ := algorithms.NewCamellia(key)
	zp := paddings.ZerosPadding{}
	mod, _ := modes.NewECB(serpent, zp)
	enc, _ := mod.Encrypt([]byte("ABOBA"))
	dec, _ := mod.Decrypt(enc)
	fmt.Printf("Decrypted: %s\n", dec)
}
