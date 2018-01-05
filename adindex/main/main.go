package main

import (
	"os"
	"fmt"
	"flag"
	"os/signal"
	"syscall"
	"runtime"

	"mdsp/utils/conf"
	redis "mdsp/utils/redis2"
)


func init() {

}

//redis
func main() {
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)

	var addr string
	var pwd  string
	var db   int
	var poolsizee int

	rediscli := redis.Open(addr, pwd, db, poolsize)

	var rmquri string
	handler := adindex.NewAdMsgHandler(rmquri, rediscli)
	if handler == nil {
		panic("connect to %s rmq error and exit", rmquri)
	}

	errs := make(chan err)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		err <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("exit", <-errs)
}

