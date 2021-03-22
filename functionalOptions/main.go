package main

import (
	"crypto/tls"
	"time"
)

type Config struct {
	Protocol string
	Timeout  time.Duration
	Maxconns int
	TLS      *tls.Config
}

type Server struct {
	*Config
	Addr string
	Port int
}

type Option func(*Server)

func protocol(p string) Option {
	return func(s *Server) {
		s.Protocol = p
	}
}

func Timeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}
func MaxConns(maxconns int) Option {
	return func(s *Server) {
		s.Maxconns = maxconns
	}
}
func TLS(tls *tls.Config) Option {
	return func(s *Server) {
		s.TLS = tls
	}
}

func NewServer(addr string, port int, options ...func(*Server)) (*Server, error) {
	srv := Server{
		Addr: addr,
		Port: port,
		Config: &Config{
			Protocol: "tcp",
			Timeout:  30 * time.Second,
			Maxconns: 1000,
			TLS:      nil,
		},
	}
	for _, option := range options {
		option(&srv)
	}

	return &srv, nil
}

func main() {
	_, _ = NewServer("/", 8080, Timeout(300*time.Second), MaxConns(1000))
}
