package models

// Message 메시지
type Message struct {
	Targets []string `json:"targets"`
	Subject string   `json:"subject"`
	Message string   `json:"message"`
}
