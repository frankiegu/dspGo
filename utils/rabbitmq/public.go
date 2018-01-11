package rabbitmq

import "fmt"
import "log"
import "github.com/streadway/amqp"


type Publisher struct {
	conn			*amqp.Connection
	channel		*amqp.Channel
	exchange			string
	exchangeType	string
}

func NewPublisher(amqpURI, exchange, exchangeType string) (*Publisher, error) {
	p := &Publisher{}

	var err error
	p.conn, err = amqp.Dial(amqpURI)
	if err != nil {
		return nil, fmt.Errorf("Dial: %s", err)
	}

	p.channel, err = p.conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("Channel: %s", err)
	}

	log.Printf("got Channel, declaring %q Exchange (%q)", exchangeType, exchange)
	if err := p.channel.ExchangeDeclare(
		exchange,     // name
		exchangeType, // type
		false,        // durable
		false,        // auto-deleted
		false,        // internal
		false,        // noWait
		nil,          // arguments
	); err != nil {
		return nil, fmt.Errorf("Exchange Declare: %s", err)
	}

	p.exchange 			= exchange
	p.exchangeType 	= exchangeType

	return p, nil
}

func (p *Publisher) Publish(body []byte, routingKey, contentType string ) error {
	log.Printf("declared Exchange, publishing %dB body (%q)", len(body), string(body))
	if err := p.channel.Publish(
		p.exchange,   // publish to an exchange
		routingKey, // routing to 0 or more queues
		false,      // mandatory
		false,      // immediate
		
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     contentType,
			ContentEncoding: "",
			Body:            body,
			DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,              // 0-9
			// a bunch of application/implementation-specific fields
		},
	); err != nil {
		return fmt.Errorf("Exchange Publish: %s", err)
	}

	return nil
}

func (p *Publisher) Close() {
	p.conn.Close()
}

