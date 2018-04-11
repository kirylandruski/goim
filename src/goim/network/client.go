package network

import (
	"context"
	"crypto/tls"
)

func Connect(ctx context.Context, address string, cert []byte) (*Connection, error) {
	var config *tls.Config
	if cert == nil {
		config = &tls.Config{InsecureSkipVerify: true}
	} else {
		config = &tls.Config{
			InsecureSkipVerify: true,
			Certificates:       []tls.Certificate{{Certificate: [][]byte{cert}}},
		}
	}

	conn, err := tls.Dial("tcp", address, config)
	if err != nil {
		return nil, err
	}

	return NewConnection(ctx, conn), nil
}
