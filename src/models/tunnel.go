package models

type Tunnel struct {
	RequestID int    `json:"request_id"`
	Data      []byte `json:"data"`
}
