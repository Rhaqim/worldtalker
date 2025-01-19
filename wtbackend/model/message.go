package model

import "time"

type Translation struct {
	Content           string `json:"content"`
	TranslatedContent string `json:"translated_content"`
	SourceLanguage    string `json:"source_language"`
	TargetLanguage    string `json:"target_language"`
}

type Message struct {
	ID         string    `json:"id"`
	SenderID   string    `json:"sender_id"`
	ReceiverID string    `json:"receiver_id"`
	Timestamp  time.Time `json:"timestamp"`
	Translation
}

func NewID() string {
	return time.Now().Format("20060102150405")
}
