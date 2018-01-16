package main


import (
	_"os"
	_"fmt"
	"flag"
	_"errors"
	"runtime"
	"mdsp/adserver"
	"mdsp/utils/conf"
	_"golang.org/x/net/context"
	"mdsp/adserver/adx"
	"mdsp/utils/log"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"fmt"
	"encoding/json"
	_"github.com/golang/protobuf/proto"
	"github.com/mxmCherry/openrtb"
)
var confile string
var adsvr *adserver.AdServer
type Config struct {
	Server	config.NetAddr
	//DB  db.DBConfig
}

func init() {
	flag.StringVar(&confile, "f", "", "-f: adserver config file")
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()
	conf := &Config{}
	if err := config.Read(confile, conf); err != nil {
		log.Logger().Error("config server error: %+v", err)
		return
	}
	adx.RegisterAction()

	r := fasthttprouter.New()
	r.POST("/dspRequest", dspRequest)
	addr := fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port)
	log.Logger().Fatal(fasthttp.ListenAndServe(addr, /*HandleRequest*/ r.Handler))
}

func dspRequest(ctx *fasthttp.RequestCtx){
	fmt.Fprintf(ctx, "Query string is %q\n", ctx.PostArgs())
	//err = proto.Unmarshal(data, openrtb.BidRequest{})
	BidValue := ctx.FormValue("BidRequest")

	//解析BidRequest
	BidRequest := &openrtb.BidRequest{}
	err := json.Unmarshal([]byte(BidValue),BidRequest)
	if (err !=  nil){
		fmt.Fprintf(ctx, "BidRequest Unmarshal error %s\n", BidValue)
		return
	}
	action,err := adx.GetAction(1)
	if (err == nil){
		fmt.Fprintf(ctx, "BidRequest %s\n", BidRequest)
		action.HandleBidding(BidRequest)
	}

	//testString := "test"
	//ctx.WriteString(testString)
}
