package adapi

import (
	"fmt"
	"errors"
	"encoding/json"
	"mdsp/common/consts"
	"mdsp/common/typedef"

	rmq "mdsp/utils/rabbitmq"
)

var (
	Publisher *rmq.Publisher
)

type UpdateCreative struct {
	Crves []PopupCrv
}
type UpdateBannerCreative struct{
	Crves []BannerCrv
}

type UpdateNativeCreative struct{
	Crves []NativeCrv
}

type UpdateRetargeting struct {
	CampId uint64
	AuIdList []string
}



type InvenMsg struct {
	CampId uint64
	InvenNames []string
}


var (
	ErrNonePublisher    = errors.New("server internal error: invalid publisher")
	ErrNotValidCampaign = errors.New("campaign lack element")
	ErrCampNativeInValid = errors.New("creative not match adtype")
)


func CreateCampaign(camp *Campaign) error {
	if camp.Basic == nil || camp.Target == nil || camp.Budget == nil {
		return ErrNotValidCampaign
	}

	switch camp.Adtype {
		case AdPopup:
			if len(camp.Popups) <= 0 {
				return ErrCampNativeInValid
			}
		case AdBanner:
			if len(camp.Banners) <= 0 {
				return ErrCampNativeInValid
			}
		case AdNative:
			if len(camp.Natives) <= 0 {
				return ErrCampNativeInValid
			}
	}

	return nil
}

func DeleteCampaign(campId uint64) error {
	if Publisher == nil {
		return ErrNonePublisher
	}

	msg := typedef.AdApi2BeMsg {
		Key	:	consts.API_KEY_DELETE_CAMPAIGN,
		Body: fmt.Sprintf("%d", campId),
	}

	msgBody, err := json.Marshal(&msg);
	if err != nil {
		return err
	}
	return Publisher.Publish([]byte(msgBody), consts.RMQ_API_ROUTINE_NAME, consts.RMQ_API_CONTENTTYPE)
}

func UpdateBasic(campId uint64, camp *Campaign) error {
	if Publisher == nil {
		return ErrNonePublisher
	}

	var basic typedef.CampBasic
	msg := typedef.AdApi2BeMsg {
		Key	:	consts.API_KEY_UPDATE_BASIC,
		Body: basic.String(),
	}
	msgBody, err := json.Marshal(&msg);
	if (err != nil) {
		return err
	}
	return Publisher.Publish([]byte(msgBody), consts.RMQ_API_ROUTINE_NAME, consts.RMQ_API_CONTENTTYPE)
}

func UpdateTarget(campId uint64, target *CampTarget) error {
	if Publisher == nil {
		return ErrNonePublisher
	}

	var tar typedef.Target
	msg := typedef.AdApi2BeMsg {
		Key	:	consts.API_KEY_UPDATE_TARGET,
		Body: tar.String(),
	}

	msgBody, err := json.Marshal(&msg);
	if (err != nil) {
		return err
	}

	return Publisher.Publish([]byte(msgBody), consts.RMQ_API_ROUTINE_NAME, consts.RMQ_API_CONTENTTYPE)
}

func UpdateBudget(campId uint64, budget *CampBudget) error {
	if Publisher == nil {
		return ErrNonePublisher
	}

	var b typedef.Budget
	msg := typedef.AdApi2BeMsg {
		Key	:	consts.API_KEY_UPDATE_BUDGET,
		Body: b.String(),
	}

	msgBody, err := json.Marshal(&msg);
	if (err != nil) {
		return err
	}

	return Publisher.Publish([]byte(msgBody), consts.RMQ_API_ROUTINE_NAME, consts.RMQ_API_CONTENTTYPE)
}

func UpdatePopup(campId uint64, popups []PopupCrv) error {
	if Publisher == nil {
		return ErrNonePublisher
	}

	upCrv := UpdateCreative {
		Crves: make([]PopupCrv, len(popups)),
	}

	for i, n := range popups {
		upCrv.Crves[i] = n
	}

	msg := typedef.AdApi2BeMsg {
		Key	:	consts.API_KEY_UPDATE_POPUP,
		Body: ,
	}

	msgBody, err := json.Marshal(&msg);
	if (err != nil) {
		return err
	}

	return Publisher.Publish([]byte(msgBody), consts.RMQ_API_ROUTINE_NAME, consts.RMQ_API_CONTENTTYPE)
}

func UpdateBanner(campId uint64, banner []BannerCrv) error {
	if Publisher == nil {
		return ErrNonePublisher
	}

	upCrv := UpdateBannerCreative {
		Crves: make([]BannerCrv, len(banner)),
	}

	for i, n := range banner {
		upCrv.Crves[i] = n
	}

	msg := typedef.AdApi2BeMsg {
		Key	:	consts.API_KEY_UPDATE_BANNER,
		Body: ,
	}

	msgBody, err := json.Marshal(&msg);
	if (err != nil) {
		return err
	}

	return Publisher.Publish([]byte(msgBody), consts.RMQ_API_ROUTINE_NAME, consts.RMQ_API_CONTENTTYPE)
}

func UpdateNative(campId uint64, native []NativeCrv ) error {
	if Publisher == nil {
		return ErrNonePublisher
	}

	upCrv := UpdateNativeCreative {
		Crves: make([]NativeCrv,len(native)),
	}

	for i, n := range native {
		upCrv.Crves[i] = n.string()
	}

	upCrvBody, err := json.Marshal(&upCrv)
	msg := typedef.AdApi2BeMsg {
		Key	:	consts.API_KEY_UPDATE_NATIVE,
		Body: string(upCrvBody),
	}

	msgBody, err := json.Marshal(&msg);
	if (err != nil) {
		return err
	}

	return Publisher.Publish([]byte(msgBody), consts.RMQ_API_ROUTINE_NAME, consts.RMQ_API_CONTENTTYPE)
}

func StartCampaign(campId uint64) error {
	if Publisher == nil {
		return ErrNonePublisher
	}

	msg := typedef.AdApi2BeMsg {
		Key	:	consts.API_KEY_START_CAMPAIGN,
		Body: fmt.Sprintf("%d", campId),
	}

	if msgBody, err := json.Marshal(&msg); err != nil {
		return err
	}

	return Publisher.Publish([]byte(msgBody), consts.RMQ_API_ROUTINE_NAME, consts.RMQ_API_CONTENTTYPE)
}

func PauseCampaign(campId uint64) error {
	if Publisher == nil {
		return ErrNonePublisher
	}

	msg := typedef.AdApi2BeMsg {
		Key	:	consts.API_KEY_PAUSE_CAMPAIGN,
		Body: fmt.Sprintf("%d", campId),
	}

	if msgBody, err := json.Marshal(&msg); err != nil {
		return err
	}

	return Publisher.Publish([]byte(msgBody), consts.RMQ_API_ROUTINE_NAME, consts.RMQ_API_CONTENTTYPE)
}

func ActiveCreative(campId uint64, crvIds []uint64) error {
	if Publisher == nil {
		return ErrNonePublisher
	}

	crv := typedef.CreativeMsg {
		CampId			: campId,
		CreativeIds	: crvIds,
	}

	crvBody, err := json.Marshal(&crv)
	if err != nil {
		return err
	}

	msg := typedef.AdApi2BeMsg {
		Key	:	consts.API_KEY_ACTIVE_CREATIVE,
		Body: string(crvBody),
	}

	if msgBody, err := json.Marshal(&msg); err != nil {
		return err
	}

	return Publisher.Publish([]byte(msgBody), consts.RMQ_API_ROUTINE_NAME, consts.RMQ_API_CONTENTTYPE)
}

func InactiveCreative(campId uint64, crvIds []uint64) error {
	if Publisher == nil {
		return ErrNonePublisher
	}

	crv := typedef.CreativeMsg {
		CampId			: campId,
		CreativeIds	: crvIds,
	}

	crvBody, err := json.Marshal(&crv)
	if err != nil {
		return err
	}

	msg := typedef.AdApi2BeMsg {
		Key	:	consts.API_KEY_INACTIVE_CREATIVE,
		Body: string(crvBody),
	}

	if msgBody, err := json.Marshal(&msg); err != nil {
		return err
	}

	return Publisher.Publish([]byte(msgBody), consts.RMQ_API_ROUTINE_NAME, consts.RMQ_API_CONTENTTYPE)
}

func ApproveCreative(campId uint64, crvIds []uint64) error {
	if Publisher == nil {
		return ErrNonePublisher
	}

	crv := typedef.CreativeMsg {
		CampId			: campId,
		CreativeIds	: crvIds,
	}

	crvBody, err := json.Marshal(&crv)
	if err != nil {
		return err
	}

	msg := typedef.AdApi2BeMsg {
		Key	:	consts.API_KEY_APPROVE_CREATIVE,
		Body: string(crvBody),
	}

	if msgBody, err := json.Marshal(&msg); err != nil {
		return err
	}

	return Publisher.Publish([]byte(msgBody), consts.RMQ_API_ROUTINE_NAME, consts.RMQ_API_CONTENTTYPE)
}

func UpdateAudience(campId uint64, au *Audience) error {
	if Publisher == nil {
		return ErrNonePublisher
	}

	auMsg := UpdateRetargeting {
		CampId		: campId,
		AuIdList	: au.AuList,
	}

	var key string
	if au.AuType == AudienceInclude {
		key = consts.API_KEY_UPDATE_RETARGETTING_INCLUDE
	} else if au.AuType == AudienceExclude {
		key = consts.API_KEY_UPDATE_RETARGETTING_EXCLUDE
	}

	msg := typedef.AdApi2BeMsg {
		Key	:	key,
		Body: fmt.Sprintf("%d", campId),
	}

	if msgBody, err := json.Marshal(&msg); err != nil {
		return err
	}

	return Publisher.Publish([]byte(msgBody), consts.RMQ_API_ROUTINE_NAME, consts.RMQ_API_CONTENTTYPE)
}

func UpdateInventory(campId uint64, inventory *Inventory) error {
	if Publisher == nil {
		return ErrNonePublisher
	}

	inven := InvenMsg {
		CampId		: campId,
		InvenNames: inventory.InvenList,
	}

	invenBody, err := json.Marshal(&inven)
	if err != nil {
		return err
	}

	var key string
	if inventory.InvenType == InventoryBlack {
		key = consts.API_KEY_UPDATE_BLACKINVENTORY
	} else if inventory.InvenType == InventoryWhite {
		key = consts.API_KEY_UPDATE_WHITEINVENTORY
	}

	msg := typedef.AdApi2BeMsg {
		Key	:	key,
		Body: invenBody,
	}

	if msgBody, err := json.Marshal(&msg); err != nil {
		return err
	}

	return Publisher.Publish([]byte(msgBody), consts.RMQ_API_ROUTINE_NAME, consts.RMQ_API_CONTENTTYPE)
}


func UpdateRetarget(campId uint64, retar *CampRetarget) error {
	if Publisher == nil {
		return ErrNonePublisher
	}

	var key string
	if retar.IsRetargettingEnable == 0{
		key = consts.API_KEY_UPDATE_RETARGETTING_INCLUDE
	} else {
		key = consts.API_KEY_UPDATE_RETARGETTING_EXCLUDE
	}

	msg := typedef.AdApi2BeMsg {
		Key	:	key,
		Body: fmt.Sprintf("%d", campId),
	}

	if msgBody, err := json.Marshal(&msg); err != nil {
		return err
	}

	return Publisher.Publish([]byte(msgBody), consts.RMQ_API_ROUTINE_NAME, consts.RMQ_API_CONTENTTYPE)
}
