package internal

import "time"

type Chat struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Messages  []Message `json:"messages"`
}

type Message struct {
	ID        string    `json:"id"`
	ChatID    string    `json:"chat_id"`
	Sender    string    `json:"sender"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	Feedback  *Feedback `json:"feedback,omitempty"`
}

type Feedback struct {
	Negative  bool      `json:"negative"`
	CreatedAt time.Time `json:"created_at"`
}
