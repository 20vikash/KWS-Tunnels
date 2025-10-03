package models

type Tunnel struct {
	RequestID string `json:"request_id"`
	Data      []byte `json:"data"`
}
