package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"jedi-team-BE-challenge/internal"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Chat endpoints
	r.HandleFunc("/chats", createChatHandler).Methods("POST")
	r.HandleFunc("/chats", listChatsHandler).Methods("GET")
	r.HandleFunc("/chats/{chatID}", getChatHandler).Methods("GET")
	r.HandleFunc("/chats/{chatID}/messages", sendMessageHandler).Methods("POST")
	r.HandleFunc("/chats/{chatID}/messages/{messageID}/feedback", feedbackHandler).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

// Handler stubs

type CreateChatRequest struct {
	UserID string `json:"user_id"`
}

// Creates a new chat session and returns it.
func createChatHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateChatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	fmt.Println("Creating new chat session")
	// print the req
	fmt.Printf("Request: %+v\n", req)
	chatID := uuid.New().String()
	fmt.Printf("Creating chat: %s\n", chatID)
	fmt.Println(req.UserID)
	chat := &internal.Chat{
		ID:        chatID,
		UserID:    req.UserID,
		Title:     "Chat " + chatID[:8],
		CreatedAt: time.Now(),
		Messages:  []internal.Message{},
	}
	internal.AddChat(chat)
	json.NewEncoder(w).Encode(chat)
}

// Lists all chat sessions for a user.
func listChatsHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	chats := internal.ListChats(userID)
	json.NewEncoder(w).Encode(chats)
}

// Retrieves a specific chat session.
func getChatHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	chat := internal.GetChat(vars["chatID"])
	if chat == nil {
		http.Error(w, "chat not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(chat)
}

type SendMessageRequest struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

func sendMessageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	chatID := vars["chatID"]
	chat := internal.GetChat(chatID)
	if chat == nil {
		http.Error(w, "chat not found", http.StatusNotFound)
		return
	}
	var req SendMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received message from %s: %s\n", req.Sender, req.Text)

	answer, found := internal.FindAnswer(req.Text)

	fmt.Printf("Found answer: %v\n", found)
	fmt.Printf("Answer: %s\n", answer)

	var response string
	if found {
		response = answer
	} else {
		response = "Sorry, I can't assist with that."
	}
	msgID := uuid.New().String()
	userMsg := internal.Message{
		ID:        msgID,
		ChatID:    chatID,
		Sender:    req.Sender,
		Text:      req.Text,
		CreatedAt: time.Now(),
	}
	internal.AddMessage(chatID, userMsg)
	botMsg := internal.Message{
		ID:        uuid.New().String(),
		ChatID:    chatID,
		Sender:    "bot",
		Text:      response,
		CreatedAt: time.Now(),
	}
	internal.AddMessage(chatID, botMsg)
	// Auto-title: use first user message
	// if len(chat.Messages) == 0 {
	// 	chat.Title = autoTitle(req.Text)
	// }
	json.NewEncoder(w).Encode(botMsg)
}

// func autoTitle(text string) string {
// 	words := strings.Fields(text)
// 	if len(words) > 6 {
// 		return strings.Join(words[:6], " ") + "..."
// 	}
// 	return text
// }

type FeedbackRequest struct {
	Negative bool `json:"negative"`
}

func feedbackHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	chatID := vars["chatID"]
	msgID := vars["messageID"]
	chat := internal.GetChat(chatID)
	if chat == nil {
		http.Error(w, "chat not found", http.StatusNotFound)
		return
	}
	var req FeedbackRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	for i := range chat.Messages {
		if chat.Messages[i].ID == msgID {
			chat.Messages[i].Feedback = &internal.Feedback{
				Negative:  req.Negative,
				CreatedAt: time.Now(),
			}
			json.NewEncoder(w).Encode(chat.Messages[i])
			return
		}
	}
	http.Error(w, "message not found", http.StatusNotFound)
}
