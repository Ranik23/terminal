# Используем базовый образ golang:alpine
FROM golang:alpine

WORKDIR /root/ssh_connections_manager

COPY . .

RUN go mod tidy

EXPOSE 8080

RUN go build -o manager cmd/main.go

CMD ["./manager"]
