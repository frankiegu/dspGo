//# dsp events service
package main

import (
	""
	"flag"

	"dsp/units/conf"
	"github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)



const (
	TRAN_ID     = "id"
	USER_ID			= "usr"
	CA_ID				= "ca"
	CRV_ID			= "crv"
	PRICE       = "price"
	URL         = "url"
	ADX         = "adx"
	CUR         = "cur"
	DEV         = "dev"
	DEVT        = "devt"
	ASYNC       = "async"
	KEY         = "key"
	LDP         = "ldp"
	TRACKING		= "trk"
	TRAFFICS		= "trs"
)

type ImpEventMsg struct {
	CaId			uint64
	UId				uint64
	Price			uint64

	CurTime		uint64
	DevId			string
	Key				string
	Place			string
	Adx				string
}

type ClickEventMsg struct {
	CaId			uint64
	UId			uint64

	CurTime		uint64
	DevId			string
	Key				string
	Place			string
	Adx				string
}

type ConversionMsg struct {
	CaId			uint64
	UId			uint64

	CurTime		uint64
	DevId			string
	Key				string
	Refer			string
	Place			string
	Adx				string
}


func HandleWinNotice(ctx *fasthttp.RequestCtx) {

}


func HandleImpression(ctx *fasthttp.RequestCtx) {

}



func HandleClick(ctx *fasthttp.RequestCtx) {

}

func HandleVisit(ctx *fasthttp.RequestCtx) {

}


func HandleConversion(ctx *fasthttp.RequestCtx) {

}

func main() {

}
