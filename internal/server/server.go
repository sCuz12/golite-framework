package server

import (
	"context"
	"net/http"
	"time"
)

type Server interface {
	Start() error
	Stop() error
}


type HTTPServer struct {
	srv *http.Server
}



func NewHTTPServer(address string, handler http.Handler) *HTTPServer {
	return &HTTPServer{
		srv: &http.Server{
			Addr: address,
			Handler: handler,
		},
	}
}


func (s *HTTPServer) Start() error {
	return s.srv.ListenAndServe();
}


func (s *HTTPServer) Stop() error{
	ctx,cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	return s.srv.Shutdown(ctx)

}