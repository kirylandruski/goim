package server

import (
	"goim/network"
	"context"
	"log"
	"goim/rpc"
)

type AppHandler struct {
	Authenticator Authenticator
	Storer        Storer
}

func (h *AppHandler) Handle(ctx context.Context, connection *network.Connection) {
	adapter := newAppServerAdapter(h.Authenticator, h.Storer)
	resolver := rpc.NewRequestResolver(adapter)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			buf, err := connection.Read()
			if err != nil {
				return
			}

			req, err := rpc.DynamicDeserialize(buf)
			if err != nil {
				log.Printf("error parsing request: %v\n", err.Error())
				connection.Close()
				return
			}

			resp, err := resolver.Resolve(req)
			if err != nil {
				log.Printf("error handing request: %v\n", err.Error())
				connection.Close()
				return
			}

			connection.Write(resp.Serialize())
		}
	}
}
