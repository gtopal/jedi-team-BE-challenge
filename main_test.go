package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func setupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/chats", createChatHandler).Methods("POST")
	r.HandleFunc("/chats", listChatsHandler).Methods("GET")
	r.HandleFunc("/chats/{chatID}", getChatHandler).Methods("GET")
	r.HandleFunc("/chats/{chatID}/messages", sendMessageHandler).Methods("POST")
	r.HandleFunc("/chats/{chatID}/messages/{messageID}/feedback", feedbackHandler).Methods("POST")
	return r
}

func TestCreateChatHandler(t *testing.T) {
	r := setupRouter()
	body := []byte(`{"user_id":"testuser"}`)
	req := httptest.NewRequest("POST", "/chats", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var resp map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("invalid json: %v", err)
	}
	if resp["user_id"] != "testuser" {
		t.Errorf("expected user_id to be testuser, got %v", resp["user_id"])
	}
}

func TestListChatsHandler(t *testing.T) {
	r := setupRouter()
	// Create a chat first
	body := []byte(`{"user_id":"testuser"}`)
	req := httptest.NewRequest("POST", "/chats", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	// List chats
	req2 := httptest.NewRequest("GET", "/chats?user_id=testuser", nil)
	w2 := httptest.NewRecorder()
	r.ServeHTTP(w2, req2)
	if w2.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w2.Code)
	}
	var resp []map[string]interface{}
	if err := json.Unmarshal(w2.Body.Bytes(), &resp); err != nil {
		t.Fatalf("invalid json: %v", err)
	}
	if len(resp) == 0 {
		t.Errorf("expected at least one chat")
	}
}

func TestGetChatHandler(t *testing.T) {
	r := setupRouter()
	// Create a chat first
	body := []byte(`{"user_id":"testuser"}`)
	req := httptest.NewRequest("POST", "/chats", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var resp map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("invalid json: %v", err)
	}
	chatID := resp["id"].(string)
	// Get chat
	req2 := httptest.NewRequest("GET", "/chats/"+chatID, nil)
	w2 := httptest.NewRecorder()
	r.ServeHTTP(w2, req2)
	if w2.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w2.Code)
	}
}

func TestSendMessageHandler(t *testing.T) {
	r := setupRouter()
	// Create a chat first
	body := []byte(`{"user_id":"testuser"}`)
	req := httptest.NewRequest("POST", "/chats", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var resp map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("invalid json: %v", err)
	}
	chatID := resp["id"].(string)
	// Send message
	msgBody := []byte(`{"sender":"testuser","text":"Gen Z in Nashville"}`)
	req2 := httptest.NewRequest("POST", "/chats/"+chatID+"/messages", bytes.NewBuffer(msgBody))
	req2.Header.Set("Content-Type", "application/json")
	w2 := httptest.NewRecorder()
	r.ServeHTTP(w2, req2)
	if w2.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w2.Code)
	}
}

func TestFeedbackHandler(t *testing.T) {
	r := setupRouter()
	// Create a chat first
	body := []byte(`{"user_id":"testuser"}`)
	req := httptest.NewRequest("POST", "/chats", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var resp map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("invalid json: %v", err)
	}
	chatID := resp["id"].(string)
	// Send message
	msgBody := []byte(`{"sender":"testuser","text":"Gen Z in Nashville"}`)
	req2 := httptest.NewRequest("POST", "/chats/"+chatID+"/messages", bytes.NewBuffer(msgBody))
	req2.Header.Set("Content-Type", "application/json")
	w2 := httptest.NewRecorder()
	r.ServeHTTP(w2, req2)
	var msgResp map[string]interface{}
	if err := json.Unmarshal(w2.Body.Bytes(), &msgResp); err != nil {
		t.Fatalf("invalid json: %v", err)
	}
	msgID := msgResp["id"].(string)
	// Give feedback
	fbBody := []byte(`{"negative":true}`)
	req3 := httptest.NewRequest("POST", "/chats/"+chatID+"/messages/"+msgID+"/feedback", bytes.NewBuffer(fbBody))
	req3.Header.Set("Content-Type", "application/json")
	w3 := httptest.NewRecorder()
	r.ServeHTTP(w3, req3)
	if w3.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w3.Code)
	}
}
