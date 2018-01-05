package dao

import (
	"log"
	"errors"

	"mdsp/common/typedef"
	"github.com/go-redis/redis"
)


/*
* active campaign
* reids key : active_campaign_ids
*
*/

func ActiveCampaign(cli *redis.Client, ids ...uint64) error {
	if cli == nil {
		return ErrRedisCliNullPtr
	}

	res := cli.SAdd(REDISKEY_ACTIVE_CAMPAIGN, ids)
	return res.Err()
}

func InactiveCampaign(cli *redis.Client, ids ...uint64) error {
	if cli == nil {
		return ErrRedisCliNullPtr
	}

	res := cli.SRem(REDISKEY_ACTIVE_CAMPAIGN, ids)
	return res.Err()
}

func GetActiveCampaign(cli *redis.Client) ([]uint64, error) {
	if cli == nil {
		return nil , ErrRedisCliNullPtr
	}

	var res []uint64
	if err := cli.SMembers(REDISKEY_ACTIVE_CAMPAIGN).ScanSlice(res); err != nil {
		return nil, err
	}

	return res, nil
}

/*
*	Inactive campaign in placement
* redis key:
* 						app  : inactive_campaign_app_bundle_%s_ids
*										 xxx : app bundle id

*							site : inactive_campaign_site_domain_%s_ids
*										 xxx: domain url
*				
*/
func ActiveAppCampaign(cli *redis.Client, bundleId string, campId []uint64) error {
	if cli == nil {
		return ErrRedisCliNullPtr
	}

	key := fmt.Sprintf(REDIS_KEY_INACTIVE_CAMPAIGN_APP_IDS, bundleId)
	res := cli.SRem(key, ids)
	return res.Err()
}

func InactiveAppCampaign(cli *redis.Client, bundleId string, campId []uint64) error {
	if cli == nil {
		return ErrRedisCliNullPtr
	}

	key := fmt.Sprintf(REDIS_KEY_INACTIVE_CAMPAIGN_APP_IDS, bundleId)
	res := cli.SAdd(key, ids)
	return res.Err()
}

func GetInactiveAppCampaign(cli *redis.Client, bundleId string) ([]uint64, error) {
	if cli == nil {
		return ErrRedisCliNullPtr
	}

	var res []uint64
	key := fmt.Sprintf(REDIS_KEY_INACTIVE_CAMPAIGN_APP_IDS, bundleId)
	if err := cli.SMembers(key).ScanSlice(res); err != nil {
		return nil, err
	}
}

func ActiveSiteCampaign(cli *redis.Client, domain string, campId []uint64) error {
	if cli == nil {
		return ErrRedisCliNullPtr
	}

	key := fmt.Sprintf(REDIS_KEY_INACTIVE_CAMPAIGN_SITE_IDS, bundleId)
	res := cli.SRem(key, ids)
	return res.Err()
}

func InactiveSiteCampaign(cli *redis.Client, domain string, campId []uint64) error {
	if cli == nil {
		return ErrRedisCliNullPtr
	}

	key := fmt.Sprintf(REDIS_KEY_INACTIVE_CAMPAIGN_APP_IDS, bundleId)
	res := cli.SAdd(key, ids)
	return res.Err()
}

func GetInactiveSiteCampaign(cli *redis.Client, domain string) ([]uint64, error) {
	if cli == nil {
		return ErrRedisCliNullPtr
	}

	var res []uint64
	key := fmt.Sprintf(REDIS_KEY_INACTIVE_CAMPAIGN_APP_IDS, bundleId)

	if err := cli.SMembers(key).ScanSlice(res); err != nil {
		return nil, err
	}

	return res, nil
}
