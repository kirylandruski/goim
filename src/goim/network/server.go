package network

import (
	"crypto/tls"
	"context"
	"net"
)

type Server struct {
	handler Handler
	config  *tls.Config
}

func NewServer(handler Handler, certFilePath string, keyFilePath string) (*Server, error) {
	s := &Server{}
	s.handler = handler

	keyPair, err := tls.LoadX509KeyPair(certFilePath, keyFilePath)
	if err != nil {
		return nil, err
	}

	s.config = &tls.Config{Certificates: []tls.Certificate{keyPair}}

	return s, nil
}

func (s *Server) Start(ctx context.Context, address string) error {
	listener, err := tls.Listen("tcp", address, s.config)

	if err != nil {
		return err
	}

	connectionsChan := make(chan net.Conn)
	done := make(chan bool)

	go func(l net.Listener) {
		for {
			c, err := l.Accept()
			if err != nil {
				done <- true
				return
			}
			connectionsChan <- c
		}
	}(listener)

	for {
		select {
		case c := <-connectionsChan:
			go s.handler.Handle(ctx, NewConnection(ctx, c))
		case <-done:
			listener.Close()
			return nil
		case <-ctx.Done():
			listener.Close()
			break
		}
	}
}
