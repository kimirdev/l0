package sub

import (
	"encoding/json"
	"l0/db"
	"l0/models"
	"l0/server"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/nats-io/stan.go"
)

var (
	clusterID = "test-cluster"
	clientID  = "test-client"
	channel   = "orders"
)

type Subscriber struct {
	connection stan.Conn
	handler    *server.Handler
}

func NewSubscriber(repos *db.Repository, handler *server.Handler) (*Subscriber, error) {
	connection, err := stan.Connect(clusterID, clientID)

	if err != nil {
		return nil, err
	}

	return &Subscriber{
		connection: connection,
		handler:    handler,
	}, nil
}

func (s *Subscriber) Subscribe() {
	s.connection.Subscribe(channel, func(msg *stan.Msg) {
		var data models.OrderDTO

		err := json.Unmarshal(msg.Data, &data)

		if err != nil {
			log.Println(err)
			return
		}

		val := validator.New()

		err = val.Struct(data)

		if err != nil {
			log.Println(err)
			return
		}

		err = s.handler.Create(data)

		if err != nil {
			log.Println(err)
			return
		}

		log.Printf("model [uid: %s] successfuly saved", data.OrderUID)
	})
}
