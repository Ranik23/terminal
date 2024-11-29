package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"ssh_connections_manager/internal/handler/local"
)

type Server struct {
	srv *http.Server
}


func NewServer() *Server {
	return &Server{}
}

func (s *Server) ApplyEndPoints() {

	http.HandleFunc("/ws", local.ServeWS)
	http.HandleFunc("/hostname", local.HostHandler)
	http.HandleFunc("/dir", local.DirHandler)
	http.HandleFunc("/username", local.UserNameHandler)


	fs := http.FileServer(http.Dir(filepath.Join(os.Getenv("HOME"), "/ssh_connections_manager/internal/static")))
	log.Println(filepath.Join(os.Getenv("HOME"), "/ssh_connections_manager/internal/static"))
	http.Handle("/", fs)
}

func (s *Server) Start(address string) error {
	s.srv = &http.Server{Addr: address}
	fmt.Printf("Сервер запущен на %s\n", address)
	return s.srv.ListenAndServe()
}


