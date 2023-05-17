package server

import (
	"time"

	"github.com/ProtonFundationGroup/seventies/v2/log"
)

type Server struct {
}

func New() (*Server, error) {
	return &Server{}, nil
}

func (s *Server) Start() error {
	for {
		log.PrintGreen("server is running ...")
		time.Sleep(time.Second)
	}
}

func (s *Server) Stop() error {
	log.PrintGreen("server has been stopped.")
	return nil
}
