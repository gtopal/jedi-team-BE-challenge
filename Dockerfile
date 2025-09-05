FROM golang:1.20-alpine
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o chatbox main.go
EXPOSE 8080
CMD ["/app/chatbox"]
