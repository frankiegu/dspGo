package confluent

import (
	"tracking/utils/log"
	"gopkg.in/confluentinc/confluent-kafka-go.v0/kafka"
)

const (
	SESSIONTIMEOUT = 60000 //ms

)

type HandleEvent func(key, value []byte) error

func NewConsumer(broker string, groupId int, topic string, buffsize int, handler HandleEvent) (*Consumer, error) {
	conf := &kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id": groupId,
		"session.timeout.ms": SESSIONTIMEOUT,
		"go.events.channel.enable": true,
		"go.application.rebalance.enable": true,
		"default.topic.config": kafka.ConfigMap{"auto.offset.reset": "latest"},
	}

	c, err := kafka.NewConsumer(conf)
	if err != nil {
		return nil, err
	}

	if err := c.Subscribe(topic,nil); err != nil {
		return nil, err
	}

	cc := &Consumer {
		Running: true,
		Handler: handler,

		Topic : topic,
		KafkaConsumer: c,
	}

	go func(c *Consumer) {
		for c.Running == true {
			ev := <-c.KafkaConsumer.Events()

			switch e := ev.(type) {
				case *kafka.Message:
					//m, _ := ev.(*kafka.Message)
					//fmt.Println(ev)
					c.Handler(e.Key, e.Value)
				case kafka.AssignedPartitions:
					log.Logger().Debugf("kafka consumer client AssignedPartitions: %v\n", e)
					c.KafkaConsumer.Assign(e.Partitions)
				case kafka.RevokedPartitions:
					c.KafkaConsumer.Unassign()
				case kafka.PartitionEOF:
					log.Logger().Infof("kafka consumer client partition eof\n")
				case kafka.Error:
					log.Logger().Errorf("kafka consumer client error: %v\n", e)
					c.Running = false
			}

		}
	}(cc)

	return cc, nil
}

type Consumer struct {
	Running bool
	Handler HandleEvent

	Topic string
	KafkaConsumer *kafka.Consumer

}

func (c *Consumer) Close() error {
	c.Running = false
	return c.KafkaConsumer.Close()
}
