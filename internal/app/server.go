package app

import (
	"fmt"
	"github.com/stetsd/micro-brick/internal/domain/services"
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
	serviceCollection := tools.Bind(services.ServiceUserName)
	router := NewHttpRouter(serviceCollection)

	server.httpServer = &http.Server{
		Addr:    net.JoinHostPort("", "8000"),
		Handler: router,
	}

	shutdownChan := make(chan error)

	go func() {
		err := server.httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			shutdownChan <- err

		}
	}()

	select {
	case err := <-shutdownChan:
		fmt.Printf("everything is gone, judges are bought!!! %v", err)
	}
	fmt.Println("START")
}

func Stop() {
	fmt.Println("STOP")
}

// TODO: listen the sys signal
