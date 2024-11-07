package server

import "net/http"

type Server struct {
	// add config here
}

func New() *Server {
	return &Server{}
}

func (s *Server) Start(addr string) error {
	// add server implementation here
	return http.ListenAndServe(addr, nil)
}
