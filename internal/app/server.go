package app

import (
	"fmt"
	"github.com/stetsd/micro-brick/internal/tools"
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
	serviceCollection := tools.Bind("user")
	router := NewHttpRouter(serviceCollection)

	server.httpServer = &http.Server{
		Addr:    net.JoinHostPort("", "8000"),
		Handler: router,
	}

	fmt.Println("START")
}

func Stop() {
	fmt.Println("STOP")
}

// TODO: listen the sys signal
