package adindex

import (
	"log"

	"common/consts"
	rmq "dsp/utils/rabbitmq"
	"github.com/go-redis/redis"
)


type AdMsgHandler struct {
	AdApiRmq				*rmq.Consumer
	AdCCRmq					*rmq.Consumer

	RedisCli				*redis.Client
}


func NewAdMsgHandler(rmquri string, cli *redis.Client) (*AdMsgHandler) {
	handler := &AdMsgHandler {
	}

	apirmq , err := rmq.NewConsumer(rmquri,
										consts.RMQ_API_EXCHANGE_NAME,
										consts.RMQ_API_EXCHANGE_TYPE,
										consts.RMQ_API_ROUTINE_NAME,
									  "",
									  h.AdApiMsgHandler,
										)
	if err != nil {
		log.Println("admedsea create ad factory rmp error:", err)
		return nil
	}

	ccrmq , err := rmq.NewConsumer(rmquri,
										consts.RMQ_CC_EXCHANGE_NAME,
										consts.RMQ_CC_EXCHANGE_TYPE,
										consts.RMQ_CC_ROUTINE_NAME,
									  "",
									  h.AdCCMsgHandler,
										)
	if err != nil {
		log.Println("admedsea create ad factory rmp error:", err)
		return nil
	}

	handler.AdApiRmq = apirmq
	handler.AdCCRmq  = ccrmq
	handler.RedisCli = cli
	return handler
}

