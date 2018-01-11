package dao

import (
	_"log"
	"errors"
	"fmt"
	"mdsp/common/typedef"
	"github.com/go-redis/redis"
	pb "github.com/golang/protobuf/proto"
)


const (
	REDISKEY_ACTIVE_CAMPAIGN	=	"active_campaign_ids"

	REDIS_KEY_INACTIVE_CAMPAIGN_APP_IDS  = "inactive_campaign_app_bundle_%s_ids"
	REDIS_KEY_INACTIVE_CAMPAIGN_SITE_IDS = "inactive_campaign_site_domain_%s_ids"

	REDISkEY_BLACKLIST = "blacklist_%s"
	REDISKEY_WHITELIST = "whitelist_%s"


	REDISKEY_CAMPAIGN_WHITELIST_SET  =  "campaign_whitelist_set_%d"
	REDISKEY_CAMPAIGN_BLACKLIST_SET  =  "campaign_blacklist_set_%d"


	REDISKEY_CAMPAIGN_INFO = "campaign_%d_info"

	/*
	* set data struct , campaign ids
	*/
)

var (
	ErrRedisCliNullPtr = errors.New("redis client null ptr")
	ErrCampaignNullPtr = errors.New("campaign null ptr")
	ErrInavlidCampaign = errors.New("invalid campaign")
	ErrInvalidCampBasic = errors.New("campaign basic null prt")
)


/*
* create campaign
* campaign info : proto string
*/
func CreateCampaign(cli *redis.Client, camp *typedef.Campaign) error {
	if cli == nil {
		return ErrRedisCliNullPtr
	}

	if camp == nil {
		return ErrCampaignNullPtr
	}

	key := fmt.Sprintf(REDISKEY_CAMPAIGN_INFO, camp.Id)
	if err := cli.Set(key, camp.String(), 0).Err(); err != nil {
		return err
	}

	if err := CreateTarget(cli, camp); err != nil {
		return err
	}

	if (camp.Target != nil && camp.Target.Autype == typedef.AudienceType_eAudienceWhite) {
		if err := CreateCampaignIncludeRetargettingList(cli, camp.Id, camp.Target.RetargetingAuListId...); err != nil {
			return err
		}
	} else if camp.Target != nil && camp.Target.Autype == typedef.AudienceType_eAudienceBlack {
		if err := CreateCampaignExcludeRetargettingList(cli, camp.Id, camp.Target.RetargetingAuListId...); err != nil {
			return err
		}
	}

	return nil
	//target , creative, black/white list, audiences retargetting, ...
}


/*
* black list of campaign
* reids key : blacklist_%s
*
*/

func IsPlacementInBlacklist(cli *redis.Client, campId uint64, placement string) (bool, error) {
	if cli == nil {
		return false,ErrRedisCliNullPtr
	}

	key := fmt.Sprintf(REDISKEY_CAMPAIGN_BLACKLIST_SET, campId)
	//这块待定，到时再看
	wlist, err := cli.SMembers(key).ScanSlice()
	if err != nil {
		return false, err
	}

	for _, list := range wlist {
		if isIn, err := cli.SIsMember(list, placement).Result(); isIn {
			return true, nil
		}
	}

	return false, nil
}

func	AppendBlacklist(cli *redis.Client, blacklistName string, placement ...string) error {
	if cli == nil {
		return ErrRedisCliNullPtr
	}

	key := fmt.Sprintf(REDISKEY_CAMPAIGN_BLACKLIST_SET, campId)
	wlistKey := fmt.Sprintf(REDISKEY_BLACKLIST, blacklistName)

	if _, err := cli.SAdd(key, wlistKey).Result(); err != nil {
		return err
	}

	cnt, err := cli.SAdd(wlistKey, placement).Result()
	if err != nil {
		return err
	}

	return nil
	//if cnt != len(placement)
}

/*
* white list of campaign
* reids key : whitelist_%s
*
*/

func IsPlacementInWhitelist(cli *redis.Client, campId uint64, placement string) (bool, error) {
	if cli == nil {
		return ErrRedisCliNullPtr
	}

	key := fmt.Sprintf(REDISKEY_CAMPAIGN_WHITELIST_SET, campId)
	wlist, err := cli.SMembers(key).ScanSlice()
	if err != nil {
		return false, err
	}

	for _, list := range wlist {
		if isIn, err := cli.SIsMember(list, placement).Result(); isIn {
			return true, nil
		}
	}

	return false, nil
}

func	AppendWhitelist(cli *redis.Client, campId uint64, whitelistName string, placement ...string) error {
	if cli == nil {
		return ErrRedisCliNullPtr
	}

	key := fmt.Sprintf(REDISKEY_CAMPAIGN_WHITELIST_SET, campId)
	wlistKey := fmt.Sprintf(REDISKEY_WHITELIST, whitelistName)

	if _, err := cli.SAdd(key, wlistKey).Result(); err != nil {
		return err
	}

	cnt, err := cli.SAdd(wlistKey, placement).Result()
	if err != nil {
		return err
	}

	return nil
	//if cnt != len(placement)
}

func UpdateCampBasic(cli *redis.Client, campId uint64, basic *typedef.Camp) error {
	if cli == nil {
		return ErrRedisCliNullPtr
	}

	if basic == nil {
		return ErrInvalidCampBasic
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

	newCamp.Basic = basic
	return cli.Set(key, newCamp.String(), 0).Err()
}

func GetCampapign(cli *redis.Client, campId uint64) (*typedef.Campaign, error) {
	if cli == nil {
		return nil, ErrRedisCliNullPtr
	}

	key := fmt.Sprintf(REDISKEY_CAMPAIGN_INFO, campId)
	serialize, err := cli.Get(key).Result()

	if err != nil {
			return nil, err
	}

	var newCamp typedef.Campaign
	if err := pb.UnmarshalText(serialize, &newCamp); err != nil {
		return nil, err
	}

	return &newCamp, nil
}

/*
*
*/

