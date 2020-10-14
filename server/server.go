package server

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/YuhriBernardes/rest_go/router"
)

type Server struct {
	instance http.Server
	router   router.Router
}

func (s *Server) GetAddress() string {
	return s.instance.Addr
}

func (s *Server) Init(port int, router router.Router) {

	s.router = router

	s.instance = http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", port),
		Handler: router.Create(),
	}

}

func (s *Server) Start() error {

	log.WithFields(log.Fields{
		"address": s.instance.Addr,
		"router":  s.router.Name(),
	}).Warn("Starting sever")

	if err := s.instance.ListenAndServe(); err != nil {
		return err
	}

	return nil

}
