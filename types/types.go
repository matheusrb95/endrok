package types

type Message struct {
	SessionID int    `json:"session_id"`
	Type      string `json:"type"`
	Data      []byte `json:"data"`
}
