package adx

import (
	_"fmt"
	_"github.com/json-iterator/go"
	_"github.com/mxmCherry/openrtb"
)

/*
* extend fields in exoclick 
*/
type ExoclickRequest struct {

}

/*func ExoclickActionNew() (*ExoclickRequest,error) {
	return &ExoclickRequest{},nil
}*/
func ExoclickActionNew()(Action,error){
	return &ExoclickRequest{},nil
}

/*func ExoclickActionRequest(id uint32) (*ExoclickRequest,error) {
	return &ExoclickRequest{},nil
}*/


