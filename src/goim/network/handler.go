package network

import "context"

type Handler interface {
	Handle(ctx context.Context, connection *Connection)
}
