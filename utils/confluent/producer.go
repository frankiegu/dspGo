package confluent

import(
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v0/kafka"
)

const (
	FLUSHTIMEOUT = 10000
)

func NewProducer( broker, topic string) (*Producer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		return nil, err
	}

	go func(p *kafka.Producer) {
		for e := range p.Events() {
			m, ok := e.(*kafka.Message)
			if ! ok {
				continue
			}

			if m.TopicPartition.Error != nil {
				fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
			} else {
				fmt.Printf("Delivered message to topic %s [%d] at offset %v\n", *m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
			}
		}
	}(p)

	pp := &Producer {
		Topic	: topic,
		KafkaProducer: p,
	}

	return pp, nil
}

type Producer struct {
	DoneChan chan bool

	Topic string
	KafkaProducer *kafka.Producer
}

func (p *Producer) Produce(key, msg []byte) error {
	kafkamsg := &kafka.Message {
		TopicPartition: kafka.TopicPartition{Topic: &p.Topic, Partition: kafka.PartitionAny},
		Key  : key,
		Value: msg,
	}

	p.KafkaProducer.ProduceChannel() <- kafkamsg
	return nil
	//return p.KafkaProducer.Produce(kafkamsg)
}

func (p *Producer) Close() {
	p.KafkaProducer.Flush(FLUSHTIMEOUT)

	p.DoneChan <- true
	p.KafkaProducer.Close()
}
