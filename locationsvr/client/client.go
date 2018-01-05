package client

import (
	"fmt"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"tracking/locationsvr/proto"
)

type LocSvrClient struct {
	Conn *grpc.ClientConn
}


func NewLocClient(addr string) *LocSvrClient {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(fmt.Sprintf("location servcie client dial addr=%s error=%v\n", addr, err))
		return nil
	}

	return &LocSvrClient {
		Conn: conn,
	}
}


func (c *LocSvrClient)Close() error {
	return c.Conn.Close()
}

func Ip2Location(ctx context.Context, conn *LocSvrClient, ip string) (*proto.IpLocation, error) {
	cli := proto.NewLocationSvrClient(conn.Conn)

	in := &proto.LocRequest {
		Ip: ip,
	}

	return cli.Ip2Location(ctx, in)
}
