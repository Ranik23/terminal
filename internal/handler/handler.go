package handler

import (
	"net/http"
	"github.com/gorilla/websocket"
)


var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type HandlerManager interface {
	ServeWS(w http.ResponseWriter, r *http.Request)
	PromptHandler(w http.ResponseWriter, r *http.Request)
	UserNameHandler(w http.ResponseWriter, r *http.Request)
	HostHandler(w http.ResponseWriter, r *http.Request)
	DirHandler(w http.ResponseWriter, r *http.Request)
}
