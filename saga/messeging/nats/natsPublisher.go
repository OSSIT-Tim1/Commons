package nats

import (
	commons "github.com/OSSIT-Tim1/Commons/saga"
	"github.com/nats-io/nats.go"
)

type Publisher struct {
	conn    *nats.EncodedConn
	subject string
}

func NewNATSPublisher(host, port, user, password, subject string) (commons., error) {
}
