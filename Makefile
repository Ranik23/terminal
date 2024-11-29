build:
	sudo docker build -t manager .

run: build
	sudo docker run -p 8080:8080 --rm manager

all:
	go run cmd/main.go
