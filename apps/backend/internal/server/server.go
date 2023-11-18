package server

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Server struct {
	httpServer *http.Server
	log        *logrus.Logger
}

func New(config Config, log *logrus.Logger, handler http.Handler) *Server {
	var server *Server
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Port),
		Handler: handler,
	}
	server = &Server{
		httpServer: httpServer,
		log:        log,
	}
	log.Info("http server was configured")
	return server
}

func (server *Server) Run() error {
	server.log.Info(fmt.Sprintf("starting http server on %v port", server.httpServer.Addr))
	return server.httpServer.ListenAndServe()
}
