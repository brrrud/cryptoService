package main

import (
	"cryptoService/contollers"
	"net/http"
)

func main() {
	//key := []byte("examplekey123456")
	//serpent, _ := algorithms.NewCamellia(key)
	//zp := paddings.ZerosPadding{}
	//mod, _ := modes.NewECB(serpent, zp)
	//enc, _ := mod.Encrypt([]byte("ABOBA"))
	//dec, _ := mod.Decrypt(enc)
	//fmt.Printf("Decrypted: %s\n", dec)

	http.HandleFunc("/loadfile", contollers.LoadFileHandler)
	http.ListenAndServe(":8090/upload", nil)
}
