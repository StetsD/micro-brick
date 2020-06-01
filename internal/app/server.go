package app

import (
	"github.com/stetsd/micro-brick/internal/domain/services"
	"github.com/stetsd/micro-brick/internal/infrastructure/logger"
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

// TODO: from env
const host = "localhost"
const port = "8000"

func Start() {
	logger.Log.Info("Configure the server")

	server := newServer()
	serviceCollection := tools.Bind(services.ServiceUserName)
	router := NewHttpRouter(&serviceCollection)

	server.httpServer = &http.Server{
		Addr:    net.JoinHostPort(host, port),
		Handler: router,
	}

	shutdownChan := make(chan error)

	go func() {
		logger.Log.Info("Server started on " + host + ":" + port)
		err := server.httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			shutdownChan <- err
		}
	}()

	select {
	case err := <-shutdownChan:
		logger.Log.Error("shutdownChan msg: " + err.Error())
	}
}

//func Stop() {
//	fmt.Println("STOP")
//}

// TODO: listen the sys signal
