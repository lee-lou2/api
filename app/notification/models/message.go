package models

// Message 메시지
type Message struct {
	Targets     []string `json:"targets"`
	Subject     string   `json:"subject"`
	Body        string   `json:"body"`
	MessageType int32    `json:"message_type"`
}
