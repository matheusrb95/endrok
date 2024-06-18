package types

type WSMessage struct {
	SessionID int    `json:"session_id"`
	Type      string `json:"type"`
	Data      []byte `json:"data"`
}
