package models

type MessagePart struct {
	MsgPartID int    `json:"msgPartID"`
	Data      string `json:"data"`
}
