package nats

import (
	messeging "github.com/OSSIT-Tim1/commons/eventSourcing/messeging"
	"github.com/nats-io/nats.go"
)

type Publisher struct {
	conn    *nats.EncodedConn
	subject string
}

func NewNatsPublisher(host, port, user, password, subject string) (messeging.Publisher, error) {
	conn, err := getConnection(host, port, user, password)
	encConn, err := nats.NewEncodedConn(conn, nats.JSON_ENCODER)
	if err != nil {
		return nil, err
	}
	return &Publisher{
		conn:    encConn,
		subject: subject,
	}, nil
}

func (p *Publisher) Publish(message interface{}) error {
	err := p.conn.Publish(p.subject, message)
	if err != nil {
		return err
	}
	return nil
}
