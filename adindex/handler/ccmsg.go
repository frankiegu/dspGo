package adindex

import (
	"log"
	"encoding/json"

	"mdsp/common/consts"
	"mdsp/common/typedef"
)


func (h AdMsgHandler) AdCCMsgHandler(body []byte) {
	if body == nil || len(body) <= 0 {
		return
	}

	log.Printf("cc2bidder msg handle\n")

	var msg typedef.AdCC2BidderMsg
	err := json.Unmarshal(body, &msg)
	if err != nil {
		log.Println("unmarshal bk msg error:", err)
		return
	}

	log.Printf("msg %+v", msg)

	switch msg.Key {
		case consts.CC_KEY_ACTIVECAMPAIGN:
			//active campaign
			ActiveCampaign(h.RedisCli,)
		case consts.CC_KEY_INACTIVECAMPAIGN:
			//inactive campaign
			InactiveCampaign(h.RedisCli, )

		default:
			log.Printf("cc2bidder msg type = %s", msg.Key)
			return
	}
}
