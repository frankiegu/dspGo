package redis

import (
	"github.com/go-redis/redis"
)

/*
* addr: "ip:port"
*/

const (
	REDIS_POOL_SIZE = 256
)

func Open(addr string, pwd string, db int, poolsize int) *redis.Client {
	if poolsize <= 0 || poolsize > REDIS_POOL_SIZE {
		poolsize = REDIS_POOL_SIZE
	}

	cli := redis.NewClient(&redis.Options {
		Addr:			addr,
		Password: pwd,
		DB:				db,
		PoolSize: poolsize,
  })

	if cli == nil {
		return nil
	}

	_ , err := cli.Ping().Result()
	if err != nil {
		panic("redis ping error: %v", err)
	}

	fmt.Printf("redis ping result: %v error: %v", pong, err)
	return cli
}

/*
func OpenFailover() *redis.Client {

}
*/

