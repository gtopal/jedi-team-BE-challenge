package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

var (
	chatsMu sync.Mutex
	chats   = map[string]*Chat{}
)

func SaveChatsToFile(filename string) error {
	chatsMu.Lock()
	defer chatsMu.Unlock()
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(chats)
}

func LoadChatsFromFile(filename string) error {
	chatsMu.Lock()
	defer chatsMu.Unlock()
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewDecoder(f).Decode(&chats)
}

func GetChat(id string) *Chat {
	chatsMu.Lock()
	defer chatsMu.Unlock()
	return chats[id]
}

func ListChats(userID string) []*Chat {
	chatsMu.Lock()
	defer chatsMu.Unlock()
	var result []*Chat
	for _, c := range chats {
		if c.UserID == userID {
			result = append(result, c)
		}
	}
	return result
}

func AddChat(chat *Chat) {
	chatsMu.Lock()
	defer chatsMu.Unlock()
	chats[chat.ID] = chat
	// print the chats
	for id, c := range chats {
		fmt.Printf("Chat ID: %s, User ID: %s, Title: %s\n", id, c.UserID, c.Title)
	}
}

func AddMessage(chatID string, msg Message) {
	chatsMu.Lock()
	defer chatsMu.Unlock()
	if c, ok := chats[chatID]; ok {
		c.Messages = append(c.Messages, msg)
	}
}
