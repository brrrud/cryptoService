package models

type Message struct {
	Key             string `json:"key"`
	CryptoAlgorithm string `json:"crypto_algorithm"`
	Padding         string `json:"padding"`
	CipherMode      string `json:"cipher_mode"`
	Content         string `json:"content"`
	Format          string `json:"format"`
	CountParts      int    `json:"countParts"`
}
