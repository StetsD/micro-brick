package app

import (
	"fmt"
	"net"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func newServer() *Server {
	server := &Server{}

	return server
}

func Start() {
	server := newServer()
	//router := NewHttpRouter()

	server.httpServer = &http.Server{
		Addr: net.JoinHostPort("", "8000"),
	}
	fmt.Println("START")
}

func Stop() {
	fmt.Println("STOP")
}

// TODO: listen the sys signal
