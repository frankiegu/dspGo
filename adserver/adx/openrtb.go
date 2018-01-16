package adx

import (
	_"github.com/json-iterator/go"
	"fmt"
	"github.com/mxmCherry/openrtb"
)

const Exoclick  = 1
const Mgid  = 2
const Smaato  = 3

type FUNC_ACTION_CREATOR func()(Action, error)
//dsp流量源列表
var dspActionList = make(map[uint32]FUNC_ACTION_CREATOR)

type AdxHandler interface {

}

type AdServer struct {

}

type AdCandidate struct{
	campId uint64
	adType uint64
	creative uint64
}

type Creative struct {
	BannerCreative int
	PopupCreative int
	NativeCreative int
}

type AdCandidates struct{
	AdCandidate []AdCandidate
	Creative Creative
}

type Action interface {
	HandleBidding(req *openrtb.BidRequest) (*AdCandidates, error)
	//RetrieveBanner(req *openrtb.BidRequest) (*AdCandidates, error)
	//RetrievePopup(req *openrtb.BidRequest) (*AdCandidates,  error)
	//RetrieveNative(req *openrtb.BidRequest) (*AdCandidates, error)
}

//注册流量源
func RegisterAction() {
	register_action(Exoclick, ExoclickActionNew)
	//register_action(Mgid, MgidActionNew,req)
	//register_action(Smaato, SmaatoActionNew,req)
}
//执行注册
func register_action(actionType uint32,creator FUNC_ACTION_CREATOR) (error) {
	if creator == nil {
		return fmt.Errorf("%d register action is nil", actionType)
	}
	if actionType == 0 {
		return fmt.Errorf("%d register action is invalid", actionType)
	}
	if _, ok := dspActionList[actionType]; ok {
		return fmt.Errorf("%d action already registered", actionType)
	} else {
		dspActionList[actionType] = creator;
		return nil
	}
}

/*
获取action
 */
func GetAction(actionType uint32) (Action,error) {
	// check
	if actionType == 0 {
		return nil,fmt.Errorf("%d get action is invalid", actionType)
	}
	// run
	if creator, ok := dspActionList[actionType]; ok {
		return creator()
	} else {
		return nil,fmt.Errorf("%d get action is nil", actionType)
	}
}