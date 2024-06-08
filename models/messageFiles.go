package models

type MessageLikeFile struct {
	Message
	Format string `json:"format"` // ==кодировка(.txt, .jpg и тд )
}
