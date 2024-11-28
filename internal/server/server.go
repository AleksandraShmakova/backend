package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	srv http.Server
}

func New(addr string) *Server {
	return &Server{
		srv: http.Server{
			Addr: addr,
		},
	}
}

func (s *Server) ListenAndServe() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from custom server!")
	})
	return s.srv.ListenAndServe()
}
