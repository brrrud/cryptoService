package models

type Message struct {
	Key             string `json:"key"`
	CryptoAlgorithm string `json:"crypto_algorithm"`
	Padding         string `json:"padding"`
	CipherMode      string `json:"cipher_mode"`
	Content         string `json:"content"` //когда загружаем файл, то это ссылка на путь к файлу,// если это просто сообщение, то это текст самого сообщения
}
