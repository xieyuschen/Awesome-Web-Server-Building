package sun

import (
	"context"
	"log"
	"net"
)

type Engine struct {
	logger   *log.Logger
	listener net.Listener
	context  context.Context
}
