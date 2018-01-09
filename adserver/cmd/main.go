package main


import (
	"os"
	"fmt"
	"flag"
	"errors"
	"runtime"


	"mdsp/adserver"
	"mdsp/utils/conf"

	"golang.org/x/net/context"
	"github.com/valyala/fasthttp"
)


var confile string
var adsvr *adserver.AdServer

func init() {
	flag.StringVar(&confile, "f", "", "-f: adserver config file")
}


func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())


	flag.Parse()
	conf := &Config{}
	if err := config.Read(confile, conf); err != nil {
		log.Fatalln("config server error: ", err)
		return
	}


	//adsvr = adserver.NewAdServer()


	addr := fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port)
	s := &fasthttp.Server{
		Handler: adsvr.HandleRequest,
		Name:    "adbund dsp server",
	}

	if err := s.ListenAndServe(addr); err != nil {
		log.Fatalln("listen and server error: ", err)
	}
}

func flow() {

}

func HandleAdx(req interface{} ) (resp interface{} , err error) {
}
