package server

import (
	"fmt"
	"log"
	"main/server/session"
	"net"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func Run(config *session.Config, handler http.Handler) {
	var server Server

	server.httpServer = &http.Server{
		Handler: handler,
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", config.Listen.Port))
	if err != nil {
		log.Fatal(err)
	}

	server.httpServer.Serve(listener)
}
