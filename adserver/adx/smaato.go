package adx

import (
	_"fmt"
	_"github.com/json-iterator/go"
	"github.com/mxmCherry/openrtb"
)


/*
* extend fields in smaato
*/
type SmaatoRequest struct {

}


type SmattoBidRequestExt struct {
	xudih				string			`json:"x_uidh,omitempty"`
	opera				int					`json:"operaminibrowser,omitempty"`
	carriername string			`json:"carriername,omitempty"`
	udi *SmattoBidReqUDI		`json:"udi,omitempty"`
}

type SmattoBidReqUDI struct {

}


/*
	required fields followed
	id :   string
	impid: string
	price: float
	nurl:  string, 

	cid : campaign id
	crid: 
	iurl:
	attr:
	adomain:
*/
type SmaatoResponse struct {
	openrtb.BidResponse
}

func SmaatoActionNew(req *openrtb.BidRequest) (error) {
	return nil
}

