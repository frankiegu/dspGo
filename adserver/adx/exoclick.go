package adx

import (
	_"fmt"
	_"github.com/json-iterator/go"
	"github.com/mxmCherry/openrtb"
	"mdsp/utils/log"
)

/*
* extend fields in exoclick 
*/
type ExoclickRequest struct {
	action_id   uint32
}

func ExoclickActionNew()(Action,error){
	return &ExoclickRequest{},nil
}

func (*ExoclickRequest)HandleBidding(req *openrtb.BidRequest) (*AdCandidates, error) {
	log.Logger().Infof("HandleBidding request %+v",req)
	
	return nil,nil
}

/*func ExoclickActionRequest(id uint32) (*ExoclickRequest,error) {
	return &ExoclickRequest{},nil
}*/


