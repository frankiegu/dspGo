package dao

import (
	"log"

	"github.com/go-redis/redis"

)

const (
	REDISKEY_RETARGETTING_VIEWER		 = "retargetting_viewer_%s"
	REDISKEY_RETARGETTING_VISITOR	   = "retargetting_visitor_%s"
	REDISKEY_RETARGETTING_CONVERSER	 = "retargetting_converser_%s"

	REDISKEY_CAMPAIGN_INCLUDE_RETARGETTING_SET	= "campaign_include_retargetting_set_%d"
	REDISKEY_CAMPAIGN_EXCLUDE_RETARGETTING_SET	= "campaign_exclude_retargetting_set_%d"

)


/*
*create : delete older memebers, and then add newer memebers
*/

func CreateCampaignIncludeRetargettingList(cli *redis.client, campid uint64, listname ...string) error {
	if cli == nil {
		return errredisclinullptr
	}

	key := fmt.sprintf(REDISKEY_CAMPAIGN_INCLUDE_RETARGETTING_SET, campid)
	exist, err :=  cli.Exists(key)
	if err != nil {
		return err
	}

	if exist == 1 {
		cli.Del(key)
	}
	return cli.SAdd(key, listname).Err()
}

func CreateCampaignExcludeRetargettingList(cli *redis.client, campid uint64, listname ...string) error {
	if cli == nil {
		return errredisclinullptr
	}

	key := fmt.sprintf(REDISKEY_CAMPAIGN_EXCLUDE_RETARGETTING_SET, campid)
	exist, err :=  cli.Exists(key)
	if err != nil {
		return err
	}

	if exist == 1 {
		cli.Del(key)
	}
	return cli.SAdd(key, listname).Err()
}


/*
* viewer retargetting list
* reids key : retargetting_viewer_%s
*
*/

func AppendDev2ViewerList(cli *redis.Client, viewerListName string, devIds ...string) error {
	if cli == nil {
		return ErrRedisCliNullPtr
	}

	viewerKey := fmt.Sprintf(REDISKEY_RETARGETTING_VIEWER, viewerListName)

	_ , err := cli.SAdd(viewerKey, devIds).Result()
	return err
}

/*
* visitor retargetting list
* reids key : retargetting_visitor_%s
*
*/

func AppendDev2VisitorRetargetting(cli *redis.Client, visitorListName string, devIds ...string) error {
	if cli == nil {
		return ErrRedisCliNullPtr
	}

	visitorKey := fmt.Sprintf(REDISKEY_RETARGETTING_VISITOR, visitorListName)

	_ , err := cli.SAdd(wlistKey, placement).Result()
	return err
}


/*
* converser retargetting list
* reids key : retargetting_converser_%s
*
*/

func AppendDev2ConverserList(cli *redis.Client, converserListName string, devIds ...string) error {
	if cli == nil {
		return ErrRedisCliNullPtr
	}

	converserKey := fmt.Sprintf(REDISKEY_RETARGETTING_CONVERSER, converserListName)

	_ , err := cli.SAdd(converserKey, devIds).Result()
	return err
}



func IsDevInRetargettingList(cli *redis.Client, campId uint64, devId string, isInclude bool) (bool, error) {
	if cli == nil {
		return ErrRedisCliNullPtr
	}

	key := fmt.Sprintf(REDISKEY_CAMPAIGN_INCLUDE_RETARGETTING_SET, campId)
	if ! isInclude {
		key = fmt.Sprintf(REDISKEY_CAMPAIGN_EXCLUDE_RETARGETTING_SET, campId)
	}

	rlist, err := cli.SMembers(key).ScanSlice()
	if err != nil {
		return false, err
	}

	for _, list := range rlist {
		if isIn, err := cli.SIsMemeber(list, devId).Result(); isIn && err == nil {
			return true, nil
		}
	}

	return false, nil
}

/*
*
*/

func IsDevInViewerList(cli *redis.Client, campId uint64, devId string, isInclude bool) (bool, error) {
	if cli == nil {
		return ErrRedisCliNullPtr
	}

	key := fmt.Sprintf(REDISKEY_CAMPAIGN_INCLUDE_RETARGETTING_SET, campId)
	if ! isInclude {
		key = fmt.Sprintf(REDISKEY_CAMPAIGN_EXCLUDE_RETARGETTING_SET, campId)
	}

	vlist, err := cli.SMembers(key).ScanSlice()
	if err != nil {
		return false, err
	}

	for _, list := range vlist {
		if isIn, err := cli.SIsMemeber(list, devId).Result(); isIn {
			return true, nil
		}
	}

	return false, nil
}


func IsDevInVisitorList(cli *redis.Client, campId uint64, devId string, isInclude bool) (bool, error) {
	if cli == nil {
		return ErrRedisCliNullPtr
	}

	key := fmt.Sprintf(REDISKEY_CAMPAIGN_INCLUDE_RETARGETTING_SET, campId)
	if ! isInclude {
		key = fmt.Sprintf(REDISKEY_CAMPAIGN_EXCLUDE_RETARGETTING_SET, campId)
	}

	vlist, err := cli.SMembers(key).ScanSlice()
	if err != nil {
		return false, err
	}

	for _, list := range vlist {
		if isIn, err := cli.SIsMemeber(list, devId).Result(); isIn {
			return true, nil
		}
	}

	return false, nil
}

//
func IsDevInConverserList(cli *redis.Client, campId uint64, devId string, isInclude bool) (bool, error) {
	if cli == nil {
		return ErrRedisCliNullPtr
	}

	key := fmt.Sprintf(REDISKEY_CAMPAIGN_INCLUDE_RETARGETTING_SET, campId)
	if ! isInclude {
		key = fmt.Sprintf(REDISKEY_CAMPAIGN_EXCLUDE_RETARGETTING_SET, campId)
	}

	clist, err := cli.SMembers(key).ScanSlice()
	if err != nil {
		return false, err
	}

	for _, list := range clist {
		if isIn, err := cli.SIsMemeber(list, devId).Result(); isIn {
			return true, nil
		}
	}

	return false, nil
}
