package rabbitmq

import (
	"log"
	"fmt"
  "github.com/streadway/amqp"
)

type MsgHandle func(body []byte) 

type Consumer struct {
	conn 			*amqp.Connection
	channel 	*amqp.Channel
	tag				string
	
	done 			chan error
	finish		chan bool
	msgHandle	MsgHandle
}

func NewConsumer(amqpURI string, exchange, exchangeType, key, ctag string, msghandle MsgHandle) (*Consumer, error) {
	c :=  &Consumer {
		conn:	nil,
		channel: nil,
		tag: ctag,

		done:   make(chan error),
		finish: make(chan bool),
		msgHandle: msghandle,
	}

	var err error
	c.conn, err  = amqp.Dial(amqpURI)
	if err != nil {
		return nil, fmt.Errorf("conn error:", err)
	}

	c.channel, err = c.conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("channel error:", err)
	}

	err = c.channel.ExchangeDeclare(
		exchange,
		exchangeType,
		false,//true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, fmt.Errorf("exchange declare error:", err)
	}

	q, err := c.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)

	if err != nil {
		return nil, fmt.Errorf("queue declare error:", err)
	}

	err = c.channel.QueueBind(
		q.Name,
		key,
		exchange,
		false,
		nil,
	)

	if err != nil {
		return nil, fmt.Errorf("queue bind error:", err)
	}

	deliveries, err := c.channel.Consume(
		q.Name,
		c.tag,		//consumerTag
		false, 	  //noack
		false,	  //exclusive
		false,		//nolocal
		false,		//nowait
		nil,			//arguments
	)

	if err != nil {
		return nil, fmt.Errorf("channel consume error:", err)
	}

	go c.handle(deliveries, c.done)
	return c , nil
}

func (c *Consumer) ShutDown() error {
	if err := c.channel.Cancel(c.tag, true); err != nil {
		return fmt.Errorf("conumser channel cancel error:", err)
	}

	if err := c.conn.Close(); err != nil {
		return fmt.Errorf("consumer connection close error:", err)
	}

	defer fmt.Println("amqp shotdown ok")

	c.finish <- true
	return <-c.done
}

/*
func (c *Consumer)handle(msgs <-chan amqp.Delivery, done chan error) {
	for {
		select {
			case <-msgs:
					for msg := range msgs {
						if c.msgHandle != nil {
							c.msgHandle(msg.Body)
						}
						msg.Ack(false)
					}
			case <-c.finish :
				log.Println("consumer handle go to finish")
				c.done <- nil
				return
		}
	}
}
*/

func (c *Consumer) handle(msgs <-chan amqp.Delivery, done chan error) {
	for m := range msgs {
		log.Printf("msg:%s", string(m.Body))
		if c.msgHandle != nil {
			c.msgHandle(m.Body)
		}
		m.Ack(false)
	}

	log.Println("deliver close channel")
	done <- nil
}
