package dao

import (
	"fmt"
	"log"
	"errors"

	"mdsp/common/typedef"
	"github.com/go-redis/redis"
)


const (
	REDIS_KEY_ACTIVE_CREATIVE	  =  "active_creative_campaign_%d_ids"
	REDIS_KEY_PENdING_ACTIVE_CREATIVE		=  "pending_active_creative_campaign_%d_ids"
	REDIS_KEY_PENdING_INACTIVE_CREATIVE	=  "pending_inactive_creative_campaign_%d_ids"
)

var (
	ErrCreativeNullPtr = errors.New("campaign creative null ptr")
	ErrCreativeInvalid = errors.New("campaign append none valid creative ")
)
/*
* active creative in campaign
* redis key : active_creative_campaign_xxx_ids
*
*/


func ActiveCreative(cli *redis.Client, campId uint64, crvIds ...uint64) error {
	if cli == nil {
		return ErrRedisCliNullPtr
	}

	key := fmt.Sprintf(REDIS_KEY_ACTIVE_CREATIVE, campId)
	activekey		:= fmt.Sprintf(REDIS_KEY_PENdING_ACTIVE_CREATIVE, campId)
	inactiveKey := fmt.Sprintf(REDIS_KEY_PENdING_INACTIVE_CREATIVE, campId)

	active := make([]uint64, 0, len(crvIds))
	pend	 := make([]uint64, 0, len(crvIds))

	for _, id := range (crvIds) {
		if cli.SIsMember(inactiveKey, id) {
			pend = append(pend, id)
		} else {
			active = append(active, id)
		}
	}

	if len(active) > 0 {
		res := cli.SAdd(key, active)
		if res.Err() != nil {
			return res.Err()
		}
	}

	if len(pend) > 0 {
		res := cli.SRem(inactivekey, pend)
		res := cli.SAdd(activeKey, pend)

		if res.Err() != nil {
			return res.Err()
		}
	}

  return nil
}

func InactiveCreative(cli *redis.Client, campId uint64,  crvIds ...uint64) error {
	if cli == nil {
		return ErrRedisCliNullPtr
	}

	key := fmt.Sprintf(REDIS_KEY_ACTIVE_CREATIVE, campId)

	activekey		:= fmt.Sprintf(REDIS_KEY_PENdING_ACTIVE_CREATIVE, campId)
	inactiveKey := fmt.Sprintf(REDIS_KEY_PENdING_INACTIVE_CREATIVE, campId)

	active := make([]uint64, 0, len(crvIds))
	pend	 := make([]uint64, 0, len(crvIds))

	for _, id := range crvIds {
		if cli.SIsMember(activekey, id) {
			pend = append(pend, id)
		} else {
			active = append(active, id)
		}
	}

	if len(active) > 0 {
		res := cli.SRem(key, active)
		if res.Err() != nil {
			return res.Err()
		}
	}

	if len(pend) > 0 {
		res := cli.SRem(activekey, pend)
		res := cli.SAdd(inactiveKey, pend)
		if res.Err() != nil {
			return res.Err()
		}
	}

  return nil
}

func GetActiveCreative(cli *redis.Client, campId uint64) (crvIds []uint64, err error) {
	if cli == nil {
		return ErrRedisCliNullPtr
	}

	key := fmt.Sprintf(REDIS_KEY_ACTIVE_CREATIVE, campId)
	if res := cli.SMembers(key); res.Err() != nil {
		return nil, res.Err()
	} else {
		ids []uint64
		if err := res.ScanSlice(ids); err != nil {
			return nil, err
		}

		return ids, nil
	}
}

func ApproveCreative(cli *redis.Client, campId uint64, crvIds ...uint64) error {
	if cli == nil {
		return ErrRedisCliNullPtr
	}

	{
		key := fmt.Sprintf(REDIS_KEY_ACTIVE_CREATIVE, campId)
		activekey		:= fmt.Sprintf(REDIS_KEY_PENdING_ACTIVE_CREATIVE, campId)
		inactiveKey := fmt.Sprintf(REDIS_KEY_PENdING_INACTIVE_CREATIVE, campId)

		active		:= make([]uint64, 0, len(crvIds))
		inactive	:= make([]uint64, 0, len(crvIds))


		for _, id := range crvIds {
			if cli.SIsMember(activekey, id).Val() {
				active = append(active, id)
			} else if cli.SIsMember(inactiveKey,id).Val() {
				inactive = append(inactive, id)
			}
		}

		if len(active) > 0 {
			res := cli.SAdd(key, active)
			res := cli.SRem(activekey, active)

			if res.Err() != nil {
				return res.Err()
			}
		}

		if len(inactive) > 0 {
			res := cli.SRem(inactivekey, inactive)
			if res.Err() != nil {
				return res.Err()
			}
		}
	}

	return nil
}

func AppendCampCreative(cli *redis.Client, campId uint64, crv *typedef.Creative) error {
	if cvr == nil {
		return ErrCreativeNullPtr
	}

	if cli == nil {
		return ErrRedisCliNullPtr
	}

	var isApproved bool
	var isActive bool
	var crvId		 uint64

	if crv.GetPopupCrv() != nil {
		isActive = crv.GetPopupCrv().IsActive
		crvId    = crv.GetPopupCrv().Id
		isApproved = crv.GetPopupCrv().IsApproved
	} else if crv.GetBannerCrv() != nil {
		crvId			= crv.GetBannerCrv().Id
		isActive	= crv.GetBannerCrv().IsActive
		isApproved = crv.GetBannerCrv().IsApproved
	} else if crv.GetNativeCrv() != nil {
		crvId			= crv.GetNativeCrv().Id
		isActive	= crv.GetNativeCrv().IsActive
		isApproved = crv.GetNativeCrv().IsApproved
	} else {
		return  ErrCreativeInvalid
	}

	if ! isApproved {
		var key string
		if isActive {
			key = fmt.Sprintf(REDIS_KEY_PENdING_ACTIVE_CREATIVE, campId)
		} else {
			key = fmt.Sprintf(REDIS_KEY_PENdING_INACTIVE_CREATIVE, campId)
		}

		res := cli.SAdd(key, crvId)
		return res.Err()
	}

	if isActive {
		key := fmt.Sprintf(REDIS_KEY_ACTIVE_CREATIVE, campId)
		res := cli.SAdd(key, crvId)
		return res.Err()
	}

	return nil
}


func UpdateCampCreative(cli *redis.Client, campId uint64, crvs []*typedef.Creative) error {
	if cli == nil {
		return ErrRedisCliNullPtr
	}

	key := fmt.Sprintf(REDISKEY_CAMPAIGN_INFO, campId)
	serialize, err := cli.Get(key).Result()

	if err != nil {
			return err
	}

	var newCamp typedef.Campaign
	if err := pb.UnmarshalText(serialize, &newCamp); err != nil {
		return err
	}

	newcrv := make([]*typedef.Creative)
	for _, cvr := range crvs {
		for i, c := range newCamp.Creatives {
			if crv.GetPopupCrv() && c.GetPopupCrv() && crv.GetPopupCrv().Id == c.GetPopupCrv().Id {
				newCamp.Creatives[i].Crv = crv
			} else {
				newcrv = append(newcrv, crv)
			}

			if crv.GetBannerCrv() && c.GetBannerCrv() && crv.GetBannerCrv().Id == c.GetBannerCrv().Id {
				newCamp.Creatives[i].Crv = crv
			} else {
				newcrv = append(newcrv, crv)
			}

			if crv.GetNativeCrv() && c.GetNativeCrv() && crv.GetNativeCrv().Id == c.GetNativeCrv().Id {
				newCamp.Creatives[i].Crv = crv
			} else {
				newcrv = append(newcrv, crv)
			}
		}
	}

	if len(newcrv) > 0 {
		newCamp.Creatives = append(newCamp.Creatives, newcrv)
	}

	//newCamp.Creatives = crvs
	cli.Set(key, newCamp.String(), 0)

	{
		 key := fmt.Sprintf(REDIS_KEY_ACTIVE_CREATIVE, campId)
		 activekey	 := fmt.Sprintf(REDIS_KEY_PENdING_ACTIVE_CREATIVE, campId)
		 inactiveKey := fmt.Sprintf(REDIS_KEY_PENdING_INACTIVE_CREATIVE, campId)

		 keys := make([]string) {key, activekey, inactiveKey}
		 cli.Del(keys)

		 for _ , crv := range crvs {
				AppendCampCreative(cli, campId, crv)
		 }
	}

	return nil
}
