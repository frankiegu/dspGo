package view

import (
	"encoding/json"
	//"errors"
	"fmt"
	//"github.com/golang/protobuf/jsonpb"
	//"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"net/http"

	"mdsp/common/consts"
	"mdsp/common/typedef"
	"mdsp/utils/rabbitmq"
)

const (
	ProcessRequestTemp    = "[view][ProcessRequest]%v"
	CreateCampTemp        = "[view][CreateCampaign]%v"
	UpdateBudgetTemp      = "[view][UpdateBudget]%v"
	UpdateCreativeTemp    = "[view][UpdateCreative]%v"
	StartCampaignTemp     = "[view][StartCampaign]%v"
	PauseCampaignTemp     = "[view][PauseCampaign]%v"
	ActiveCreativeTemp    = "[view][ActiveCreative]%v"
	InactiveCreativeTemp  = "[view][InactiveCreative]%v"
	UpdateWhiteListTemp   = "[view][UpdateWhiteList]%v"
	UpdateBlackListTemp   = "[view][UpdateBlackList]%v"
	UpdateRetargetingTemp = "[view][UpdateRetargeting]%v"

	publishMsgTemp = "[publishMsg]%v"

	/*
		assignBudgetTemp         = "[assignBudgetTemp]%v"
		assignTargetTemp         = "[assignTargetTemp]%v"
		assignCampaignTemp       = "[assignCampaignTemp]%v"
		assignCampBasicTemp      = "[assignCampBasicTemp]%v"
		assignSnippetTemp        = "[assignSnippetTemp]%v"
		assignBannerCreativeTemp = "[assignBannerCreativeTemp]%v"
		assignPopupCreativeTemp  = "[assignPopupCreativeTemp]%v"
		assignNativeCreativeTemp = "[assignNativeCreativeTemp]%v"
	*/
)

/*
type CampMsg struct {
	Id   uint64 `json:"camId"`
	Body string `json:"body"`
}
*/

type RequestPackage struct {
	FuncName string `json:"func"`
	Param    string `json:"param"`
}

type RabPublisher struct {
	Publisher   *rabbitmq.Publisher
	RoutineKey  string
	ContentType string
}

var (
	pub RabPublisher
)

func (p *RabPublisher) Publish(body []byte) (err error) {
	err = p.Publisher.Publish(body, p.RoutineKey, p.ContentType)
	if err != nil {
		log.Printf("[rabbitmq publish] %v", string(body))
	}
	return
}

func logBadRequest(w http.ResponseWriter, errMsg string) {
	log.Print(errMsg)
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(errMsg)) // open on dev env, close on product env
}

func publishMsg(w http.ResponseWriter, key, body string) (err error) {
	pubMsg := typedef.AdApi2BeMsg{
		Key:  key,
		Body: body,
	}

	pubByte, err := json.Marshal(pubMsg)
	if err != nil {
		err = fmt.Errorf(publishMsgTemp, err)
		logBadRequest(w, err.Error())
		return
	}

	err = pub.Publish(pubByte)
	if err != nil {
		err = fmt.Errorf(publishMsgTemp, err)
		logBadRequest(w, err.Error())
		return
	}

	return
}

func init() {
	var err error
	pub.Publisher, err = rabbitmq.NewPublisher(
		"amqp://guest:guest@localhost:5672", // replaced with config
		consts.RMQ_API_EXCHANGE_NAME,
		consts.RMQ_API_EXCHANGE_TYPE,
	)
	if err != nil {
		log.Panicf("rabbitmq client create error: %v\n", err)
	}

	pub.RoutineKey = consts.RMQ_API_ROUTINE_NAME
	pub.ContentType = consts.RMQ_API_CONTENTTYPE
}

// for test
func ProcessRequest2(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	s := string(b)

	log.Println("---method:", r.Method)
	log.Println("---url:", r.URL)
	log.Println("---body:", s)

	fmt.Fprintf(w, "method: %v\n", r.Method)
	fmt.Fprintf(w, "url: %v\n", r.URL)
	fmt.Fprintf(w, "body: %v\n", s)
	w.Write([]byte("hell boy"))
	w.Write([]byte("hell boy2"))
}

func ProcessRequest(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf(ProcessRequestTemp, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var pack RequestPackage
	err = json.Unmarshal(body, &pack)
	if err != nil {
		log.Printf(ProcessRequestTemp, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch pack.FuncName {
	case consts.API_KEY_CREATE_CAMPAIGN: // "CreateCampaign"
		err = CreateCampaign(w, pack.Param)
	case consts.API_KEY_UPDATE_BASIC:
		//
	case consts.API_KEY_UPDATE_BUDGET: // UpdateBudget
		err = UpdateBudget(w, pack.Param)
	case consts.API_KEY_UPDATE_TARGET:
		//
	case consts.API_KEY_UPDATE_CREATIVE: // UpdateCreative
		err = UpdateCreative(w, pack.Param)
	case consts.API_KEY_START_CAMPAIGN:
		err = SwitchCampaign(w, pack.Param, true)
	case consts.API_KEY_PAUSE_CAMPAIGN:
		err = SwitchCampaign(w, pack.Param, false)
	case consts.API_KEY_ACTIVE_CREATIVE:
		err = SwitchCreative(w, pack.Param, true)
	case consts.API_KEY_INACTIVE_CREATIVE:
		err = SwitchCreative(w, pack.Param, false)
	case consts.API_KEY_UPDATE_BLACKLIST:
		err = UpdateInventory(w, pack.Param, true)
	case consts.API_KEY_UPDATE_WHITELIST:
		err = UpdateInventory(w, pack.Param, false)
	case consts.API_KEY_UPDATE_RETARGETTING:
		err = UpdateRetarget(w, pack.Param)
	default:
		err = fmt.Errorf("func name %v is not matched", pack.FuncName)
	}

	if err != nil {
		err = fmt.Errorf(ProcessRequestTemp, err)
		logBadRequest(w, err.Error())
	}
}

func CreateCampaign(w http.ResponseWriter, param string) (err error) {
	var jcam Campaign
	err = json.Unmarshal([]byte(param), &jcam)
	if err != nil {
		err = fmt.Errorf(CreateCampTemp, err)
		logBadRequest(w, err.Error())
		return
	}

	if jcam.Id == 0 {
		err = fmt.Errorf(CreateCampTemp, "camp id is missing")
		logBadRequest(w, err.Error())
		return
	}

	var pcam typedef.Campaign
	err = assignCampaign(&jcam, &pcam)
	if err != nil {
		err = fmt.Errorf(CreateCampTemp, err)
		logBadRequest(w, err.Error())
		return
	}

	pubMsg := typedef.AdApi2BeMsg{
		Key:  consts.API_KEY_CREATE_CAMPAIGN,
		Body: pcam.String(),
	}

	pubByte, err := json.Marshal(pubMsg)
	if err != nil {
		err = fmt.Errorf(CreateCampTemp, err)
		logBadRequest(w, err.Error())
		return
	}

	err = pub.Publish(pubByte)
	if err != nil {
		err = fmt.Errorf(CreateCampTemp, err)
		logBadRequest(w, err.Error())
		return
	}

	return
}

func UpdateBudget(w http.ResponseWriter, param string) (err error) {
	var jcam Campaign
	err = json.Unmarshal([]byte(param), &jcam)
	if err != nil {
		err = fmt.Errorf(UpdateBudgetTemp, err)
		logBadRequest(w, err.Error())
		return
	}

	if jcam.Id == 0 {
		err = fmt.Errorf(UpdateBudgetTemp, "camp id is missing")
		logBadRequest(w, err.Error())
		return
	}

	if jcam.Budget == nil {
		err = fmt.Errorf(UpdateBudgetTemp, "budget is missing")
		logBadRequest(w, err.Error())
		return
	}

	var pBudget typedef.Budget
	err = assignBudget(jcam.Budget, &pBudget)
	if err != nil {
		err = fmt.Errorf(UpdateBudgetTemp, err)
		logBadRequest(w, err.Error())
		return
	}

	var pcam typedef.Campaign
	pcam.Id = jcam.Id
	pcam.Budget = &pBudget

	pubMsg := typedef.AdApi2BeMsg{
		Key:  consts.API_KEY_UPDATE_BUDGET,
		Body: pcam.String(),
	}

	pubByte, err := json.Marshal(pubMsg)
	if err != nil {
		err = fmt.Errorf(UpdateBudgetTemp, err)
		logBadRequest(w, err.Error())
		return
	}

	err = pub.Publish(pubByte)
	if err != nil {
		err = fmt.Errorf(UpdateBudgetTemp, err)
		logBadRequest(w, err.Error())
		return
	}

	return
}

func UpdateCreative(w http.ResponseWriter, param string) (err error) {
	var jcam Campaign
	err = json.Unmarshal([]byte(param), &jcam)
	if err != nil {
		err = fmt.Errorf(UpdateCreativeTemp, err)
		logBadRequest(w, err.Error())
		return
	}

	if jcam.Id == 0 {
		err = fmt.Errorf(UpdateCreativeTemp, "camp id is missing")
		logBadRequest(w, err.Error())
		return
	}

	// Creatives can be empty

	var pcam typedef.Campaign
	err = assignCampaign(&jcam, &pcam)
	if err != nil {
		err = fmt.Errorf(UpdateCreativeTemp, err)
		logBadRequest(w, err.Error())
		return
	}

	pubMsg := typedef.AdApi2BeMsg{
		Key:  consts.API_KEY_UPDATE_CREATIVE,
		Body: pcam.String(),
	}

	pubByte, err := json.Marshal(pubMsg)
	if err != nil {
		err = fmt.Errorf(UpdateCreativeTemp, err)
		logBadRequest(w, err.Error())
		return
	}

	err = pub.Publish(pubByte)
	if err != nil {
		err = fmt.Errorf(UpdateCreativeTemp, err)
		logBadRequest(w, err.Error())
		return
	}

	return
}

func SwitchCampaign(w http.ResponseWriter, param string, isStart bool) (err error) {
	var msgTemp string
	var pubKey string

	if isStart == true {
		msgTemp = StartCampaignTemp
		pubKey = consts.API_KEY_START_CAMPAIGN
	} else {
		msgTemp = PauseCampaignTemp
		pubKey = consts.API_KEY_PAUSE_CAMPAIGN
	}

	//var errInfo string
	if param == "" {
		err = fmt.Errorf(msgTemp, "id is missing")
		logBadRequest(w, err.Error())
		return
	}

	/*
		var msg typedef.CampMsg
		err = json.Unmarshal(param, &msg)
		if err != nil {
			err = fmt.Errorf(msgTemp, err)
			logBadRequest(w, err.Error())
			return
		}

		if msg.Id == 0 {
			err = fmt.Errorf(msgTemp, "id is missing")
			logBadRequest(w, err.Error())
			return
		}
	*/

	pubMsg := typedef.AdApi2BeMsg{
		Key:  pubKey,
		Body: param,
	}

	pubByte, err := json.Marshal(pubMsg)
	if err != nil {
		err = fmt.Errorf(msgTemp, err)
		logBadRequest(w, err.Error())
		return
	}

	err = pub.Publish(pubByte)
	if err != nil {
		err = fmt.Errorf(msgTemp, err)
		logBadRequest(w, err.Error())
		return
	}

	return
}

func SwitchCreative(w http.ResponseWriter, param string, isActive bool) (err error) {
	var msgTemp string
	var pubKey string

	// check
	if isActive == true {
		msgTemp = ActiveCreativeTemp
		pubKey = consts.API_KEY_ACTIVE_CREATIVE
	} else {
		msgTemp = InactiveCreativeTemp
		pubKey = consts.API_KEY_INACTIVE_CREATIVE
	}

	var cmsg typedef.CreativeMsg
	err = json.Unmarshal([]byte(param), &cmsg)
	if err != nil {
		err = fmt.Errorf(msgTemp, err)
		logBadRequest(w, err.Error())
		return
	}

	if cmsg.CampId == 0 {
		err = fmt.Errorf(msgTemp, "camp_id is missing")
		logBadRequest(w, err.Error())
		return
	}

	if cmsg.CreativeIds == nil || len(cmsg.CreativeIds) == 0 {
		err = fmt.Errorf(msgTemp, "creative_ids is missing")
		logBadRequest(w, err.Error())
		return
	}

	// publish msg
	err = publishMsg(w, pubKey, param)
	if err != nil {
		err = fmt.Errorf(msgTemp, err)
		logBadRequest(w, err.Error())
		return
	}
	return
}

// isBlack, true: black, false: white
func UpdateInventory(w http.ResponseWriter, param string, isBlack bool) (err error) {
	var msgTemp string
	var pubKey string

	// check
	if isBlack == true {
		msgTemp = UpdateBlackListTemp
		pubKey = consts.API_KEY_UPDATE_BLACKLIST
	} else {
		msgTemp = UpdateWhiteListTemp
		pubKey = consts.API_KEY_UPDATE_WHITELIST
	}

	var imsg typedef.InvenMsg
	err = json.Unmarshal([]byte(param), &imsg)
	if err != nil {
		err = fmt.Errorf(msgTemp, err)
		logBadRequest(w, err.Error())
		return
	}

	if imsg.CampId == 0 {
		err = fmt.Errorf(msgTemp, "campId is missing")
		logBadRequest(w, err.Error())
		return
	}

	// can't be nil, can be empty slice
	if imsg.InvenNames == nil {
		err = fmt.Errorf(msgTemp, "invenNames is missing")
		logBadRequest(w, err.Error())
		return
	}

	// publish msg
	err = publishMsg(w, pubKey, param)
	if err != nil {
		err = fmt.Errorf(msgTemp, err)
		logBadRequest(w, err.Error())
		return
	}
	return
}

func UpdateRetarget(w http.ResponseWriter, param string) (err error) {
	var msg typedef.UpdateRetargeting
	err = json.Unmarshal([]byte(param), &msg)
	if err != nil {
		err = fmt.Errorf(UpdateRetargetingTemp, err)
		logBadRequest(w, err.Error())
		return
	}

	if msg.CampId == 0 {
		err = fmt.Errorf(UpdateRetargetingTemp, "campId is missing")
		logBadRequest(w, err.Error())
		return
	}

	// publish msg
	err = publishMsg(w, consts.API_KEY_UPDATE_RETARGETTING, param)
	if err != nil {
		err = fmt.Errorf(UpdateRetargetingTemp, err)
		logBadRequest(w, err.Error())
		return
	}

	return
}
