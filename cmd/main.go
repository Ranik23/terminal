package main

import "ssh_connections_manager/internal/server"



func main() {
	srv := server.NewServer()

	srv.ApplyEndPoints()

	srv.Start("localhost:8080")
}