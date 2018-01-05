package adindex

import (
	"log"
	"encoding/json"

	"dsp/common/consts"
	"dsp/common/typedef"

	pb "github.com/golang/protobuf/proto"
)


func AdApiMsgHandler(body []byte) {
	if body == nil || len(body) <= 0 {
		return
	}

	log.Printf("cc2bidder msg handle\n")

	var msg typedef.AdApi2BeMsg
	err := json.Unmarshal(body, &msg)
	if err != nil {
		log.Println("unmarshal bk msg error:", err)
		return
	}

	log.Printf("msg %+v", msg)

	switch msg.Key {
		case consts.API_KEY_CREATE_CAMPAIGN:
			//create campaign
			CreateCampaign(h.RedisCli, camp)
		case consts.API_KEY_UPDATE_BASIC:
			//update campaign basic
			return
	  case consts.API_KEY_UPDATE_BUDGET:
			return
		case consts.API_KEY_UPDATE_TARGET:
			//update target
			UpdateTarget(h.RedisCli, target)
		case consts.API_KEY_UPDATE_CREATIVE:
			//update creative
		case consts.API_KEY_START_CAMPAIGN:
			//start campaign
		case consts.API_KEY_PAUSE_CAMPAIGN:
			//pause campaign
		case consts.API_KEY_ACTIVE_CREATIVE:
			//active creative	
		case consts.API_KEY_INACTIVE_CREATIVE:
			//inactive creative
		case consts.API_KEY_UPDATE_BLACKLIST:
			//update black white list
		case consts.API_KEY_UPDATE_WHITELIST:
			UpdateWhiteList(caId, []byte(wl))
		case consts.API_KEY_UPDATE_RETARGETTING:
			//update retargetting
		default:
			log.Printf("unknow msg type: %s\n", msg.Key)
			return
	}

}

