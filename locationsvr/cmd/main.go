package main

import (
	//"fmt"
	"net"
	"flag"
	"time"

	"tracking/locationsvr"
	"tracking/utils/log"
	"tracking/utils/config"

	"tracking/utils/ip2location"
	pb "tracking/locationsvr/proto"
	"google.golang.org/grpc"
)


type Config struct {
	Server string
	Iplib	string
}

var confile string

func init() {
	flag.StringVar(&confile, "f", "", "-f: adretrieval config file")
}

func main() {
	flag.Parse()
	conf := &Config{}

	if err := config.Read(confile, conf); err != nil {
		panic(err)
	}

	log.Logger().Infof("config = %+v\n", conf)

	if err := ip2location.Open(conf.Iplib); err != nil {
		log.Logger().Errorf("ip2location open %s error=%v\n", conf.Iplib, err)
		panic(err)
	}

	log.Logger().Infof("listen and server start at %s, %v",conf.Server, time.Now())

	//addr := fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port)
	lis , err := net.Listen("tcp", conf.Server)
	if err != nil {
		panic(err)
	}

	locsvr := locationsvr.NewLocSvr()

	//var opt grpc.ServerOption
	grpcsvr := grpc.NewServer()
	pb.RegisterLocationSvrServer(grpcsvr, locsvr)

	if err := grpcsvr.Serve(lis); err != nil {
		log.Logger().Fatalf("server start error: %v\n", err)
	}
}
