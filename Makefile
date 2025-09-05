build:
	go build -o chatbox main.go

run:
	go run main.go

docker-build:
	docker build -t chatbox .

docker-run:
	docker run -p 8080:8080 chatbox
