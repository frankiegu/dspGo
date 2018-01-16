package adx

import (
	_"fmt"

	_"github.com/json-iterator/go"
	_"github.com/mxmCherry/openrtb"
	"fmt"
)

const Exoclick  = 1
const Mgid  = 2
const Smaato  = 3

type FUNC_ACTION_CREATOR func()(Action, error)
//dsp流量源列表
var dspActionList = make(map[uint32]FUNC_ACTION_CREATOR)

type Action interface {

}
//注册流量源
func RegisterAction() {
	register_action(Exoclick, ExoclickActionNew)
	//register_action(Mgid, MgidActionNew,req)
	//register_action(Smaato, SmaatoActionNew,req)
}
//执行注册
func register_action(actionType uint32,creator FUNC_ACTION_CREATOR) (interface{},error) {
	if creator == nil {
		return nil,fmt.Errorf("%d creator is nil", actionType)
	}
	if actionType == 0 {
		return nil,fmt.Errorf("%d creator is invalid", actionType)
	}
	if _, ok := dspActionList[actionType]; ok {
		return nil,fmt.Errorf("%d already registered", actionType)
	} else {
		dspActionList[actionType] = creator
		return dspActionList[actionType],nil
	}
}

func GetAction(actionType uint32) (Action,error) {
	// check
	if actionType == 0 {
		return nil,fmt.Errorf("%d creator is invalid", actionType)
	}
	// run
	var action Action
	action, _ = dspActionList[actionType]
	if action == nil {
		return nil,fmt.Errorf("%d creator is nil", actionType)
	}
	return  action,nil

}