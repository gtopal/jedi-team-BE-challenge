# GWI - Jedi Team - Backend Engineering Challenge

Welcome to the engineering challenge for the Jedi Team at GWI!

This task is designed to help us understand how you approach software engineering problems and apply your skills in a real-world-inspired scenario. It focuses on backend engineering using **Go**, with optional extensions into **AI/LLMs**, **product thinking**, and **system design**. The Jedi team mainly works on and evolves the AI infrastructure of the company, so this exercise has a strong focus on that.

While the base functionality is straightforward, we encourage you to go beyond the minimum requirements — creativity, thoughtful design, and clean code are all appreciated.

## 🧪 Core Requirements

You are going to create a **chatbot** that helps GWI's clients answer questions based on market research data. Another tool has converted GWI's data into a **natural language** format and stored it in a database. You can find the data in `data.md`. You should use this data to answer users' questions.

Build a web server in **Go** that exposes this chat functionality (you decide the communication method and the necessary endpoints). The discussion within the chat should be persisted, and the user should be able to continue the conversation from where it was left off. A single user can open multiple chats.

## 🌟 Optional Enhancements

- If the answer to the user's question is not found in the data, the chatbot should decline to answer.
- The user can give negative feedback on a message.
- The chat should have an auto-generated title.
- Include a **Dockerfile** and a **Makefile** or **Taskfile** to simplify local development.
- Explain in the README how to run the application and the assumptions you made.

# GWI Chatbot Backend

This project is a Go web server that provides a chatbot for GWI's clients to answer questions based on market research data. The data is stored in `data.md` in natural language format.

## Features
- REST API for chat functionality (create chat, send message, get chat history, list chats)
- Answers are generated from `data.md` only
- Chat history is persisted (in-memory, with file save/load support)
- Multiple chats per user
- Resume conversations
- Auto-generated chat titles
- Declines to answer if info not found in data
- Negative feedback on messages
- Dockerfile and Makefile for easy development

## API Endpoints

### Create a chat
`POST /chats`
```json
{
  "user_id": "user123"
}
```
Response: Chat object

### List chats for a user
`GET /chats?user_id=user123`
Response: Array of chats

### Get chat history
`GET /chats/{chatID}`
Response: Chat object

### Send a message
`POST /chats/{chatID}/messages`
```json
{
  "sender": "user123",
  "text": "What percent of Gen Z in Nashville use Instagram?"
}
```
Response: Bot message

### Give feedback on a message
`POST /chats/{chatID}/messages/{messageID}/feedback`
```json
{
  "negative": true
}
```
Response: Message object

## Running Locally

### With Makefile
```bash
make build
make run
```

### With Docker
```bash
make docker-build
make docker-run
```

## Assumptions
- Data is only sourced from `data.md` and answers are matched by substring search.
- Persistence is in-memory for simplicity, but can be extended to file or DB.
- Auto-title uses the first user message.
- Feedback is only negative/positive.
- No authentication is implemented.

## Extending
- To use a real database, replace the storage logic in `internal/storage.go`.
- To improve answer quality, enhance the search in `internal/answer.go`.

## 🧩 Submission

Just fork the current repository and send it to us!

Good luck, potential colleague!
