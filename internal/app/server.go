package app

import (
	"github.com/stetsd/micro-brick/config"
	"github.com/stetsd/micro-brick/internal/domain/services"
	"github.com/stetsd/micro-brick/internal/infrastructure/dbDriver"
	"github.com/stetsd/micro-brick/internal/infrastructure/logger"
	"github.com/stetsd/micro-brick/internal/tools"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	config     config.Config
	httpServer *http.Server
}

func NewServer(config config.Config) *Server {
	server := &Server{config: config}

	return server
}

func (server *Server) Start() {
	logger.Log.Info("Configure the server")

	dbD, err := dbDriver.NewDbDriver(server.config)

	if err != nil {
		panic(err)
	}

	serviceCollection := tools.Bind(dbD, services.ServiceUserName)
	router := NewHttpRouter(&serviceCollection)

	server.httpServer = &http.Server{
		Addr:    net.JoinHostPort(server.config.Get(config.AppHost), server.config.Get(config.AppPort)),
		Handler: router,
	}

	shutdownChan := make(chan error)
	stopChan := make(chan os.Signal, 1)

	signal.Notify(stopChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logger.Log.Info("Server started on " + server.config.Get(config.AppHost) + ":" + server.config.Get(config.AppPort))
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
