package conduit

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// Listener is the interface that needs to be implmented by listeners
type Listener interface {
	Listen(chan *Message) error
}

// Message is the message in the queue topic
type Message struct {
	Data []byte
}

// Dialer interface defines the object that works with Dial function
type Dialer interface {
	Dial(string) (*amqp.Connection, error)
}

// ListenerOptions defines the attributes needed for the listener
type ListenerOptions struct {
	Host     string
	Port     string
	UserName string
	Password string
	Dialer   Dialer
}

// NewListener creates a new listener
func NewListener(options *ListenerOptions) Listener {
	return &rmqListener{options}
}

type rmqListener struct {
	options *ListenerOptions
}

func (l *rmqListener) Listen(msgCh chan *Message) error {

	dialString := fmt.Sprintf("amqp://%s:%s@%s:%s//", l.options.UserName, l.options.Password, l.options.Host, l.options.Port)

	conn, err := amqp.Dial(dialString)
	if err != nil {
		err = fmt.Errorf("unable to dial: %w", err)
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		err = fmt.Errorf("unable to create channel: %w", err)
		return err
	}
	msgs, err := ch.Consume(
		"status",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			fmt.Printf("Received Message: %s\n", d.Body)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

	return nil
}
