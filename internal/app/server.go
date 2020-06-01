package app

import (
	"github.com/stetsd/micro-brick/internal/domain/services"
	"github.com/stetsd/micro-brick/internal/infrastructure/logger"
	"github.com/stetsd/micro-brick/internal/tools"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	httpServer *http.Server
}

func NewServer() *Server {
	server := &Server{}

	return server
}

const host = ""

func (server *Server) Start(port string) {
	logger.Log.Info("Configure the server")

	if port == "" {
		logger.Log.Fatal("missed server port flag")
	}

	serviceCollection := tools.Bind(services.ServiceUserName)
	router := NewHttpRouter(&serviceCollection)

	server.httpServer = &http.Server{
		Addr:    net.JoinHostPort(host, port),
		Handler: router,
	}

	shutdownChan := make(chan error)
	stopChan := make(chan os.Signal, 1)

	signal.Notify(stopChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

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
	case sig := <-stopChan:
		server.Stop()
		logger.Log.Info("stopChan msg: Server was " + sig.String())
	}
}

func (server *Server) Stop() {
	err := server.httpServer.Close()
	if err != nil {
		logger.Log.Fatal(err.Error())
	}
}
