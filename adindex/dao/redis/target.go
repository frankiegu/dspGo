package dao

import (
	"log"
	"errors"

	"mdsp/common/typedef"
	"github.com/go-redis/redis"
	pb "github.com/golang/protobuf/proto"
)

var (
	ErrTargetNullPtr = errors.New("campaign target null ptr")
)

/*
*-----------------------target condition-------------------
* ad type: banner, popup, native
* adx
* carrier
* connection type
* device type
* geo: country, region, city
* --idfa/gaid--
* osv
* source type: web, app
* mime: 
*----------------------------------------------------------- 
*/

const (
	ADX				=	"adx"
	OSV				= "osv"

	CITY			= "city"
	REGION		= "region"
	COUNTRY		= "country"

	ADTYPE		= "adtype"
	DEVTYPE		= "devtype"
	CONNTYPE  = "conntype"

	SOURCETYPE = "sourcetype"
)

const (
	REDISKEY_TARGET_PREFIX		= "campaign_target_"
	REDISKEY_CAMPAIGN_TARGET	= "campaign_%d_target_key"
)

/*
* adtype=&adx=11&conntype=&devtyep=&geo=&sourcetype=&osv=
* adtype=&adx=&conntype=&devtype=&city&sourcetype&osv
*/

type SubTarget struct {
		adx	uint64
		osv	string
		geo	string

		adtype		int32
		devtype		int32
		conntype	int32

		sourcetype	int32
}

func (t *SubTarget) Key() string {
	return fmt.Sprintf("%s%s=%d&%s=%d&%s=%s&%s=%s&%s=%d&%s=%d&%s=%d",REDISKEY_TARGET_PREFIX,
																													 ADTYPE, t.adtype, ADX, t.adx, OSV, t.osv,CITY, t.geo,
																													 DEVTYPE, t.devtype, CONNTYPE, t.conntype,
																												   SOURCETYPE, t.sourcetype)
}

func CreateTarget(cli *redis.Client, camp *typedef.Campaign) error {
	if camp == nil {
		return ErrCampaignNullPtr
	}

	if camp.Basic == nil || camp.Target == nil {
		return ErrInavlidCampaign
	}

	if cli == nil {
		return ErrRedisCliNullPtr
	}

	subtargets := make([]SubTarget)
	for _, adx := range camp.Target.Adxs {
		s := SubTarget {
			adx: adx,
		}

		subtargets = append(s)
	}

	precnt = len(subtargets)
	var temp []SubTarget
	copy(temp, subtargets)

	for i := len(camp.Target.Osv); i > 1; i-- {
		subtargets = append(subtargets, temp)
	}

	for i, s := range subtargets {
		s.osv = camp.Target.Osv[i / precnt]
	}

	precnt = len(subtargets)
	//temp = subtarget
	for i := len(camp.Target.Devtype); i > 1; i-- {
		subtargets = append(subtargets, subtargets[:precnt])
	}

	for i, s := range subtargets {
		s.devtype = int32(camp.Target.Devtype[i / precnt]
	}

	precnt = len(subtargets)

	connTypes := make([]int32, 0, 2)
	if camp.Target.Conntype == typedef.ConnType_eConnAll {
		connTypes = append(connTypes, int32(typedef.ConnType_eConnWifi))
		connTypes = append(connTypes, int32(typedef.ConnType_eConnMobile))
	} else {
		connTypes = append(connTypes, int32(camp.Target.Conntype)
	}

	for i := len(connTypes); i > 1; i-- {
		subtargets = append(subtargets, subtargets[:precnt])
	}

	for i, s := range subtargets {
		s.conntype = connTypes[i / precnt]
	}

	precnt = len(subtargets)

	srcTypes := make([]int32, 0, 2)
	if camp.Target.Srctype == typedef.SourceType_eSourceAll {
		srcTypes = append(srcTypes, int32(typedef.SourceType_eSourceInApp))
		srcTypes = append(srcTypes, int32(typedef.SourceType_eSourceWeb))
	} else {
		srcTypes = append(srcTypes, int32(camp.Target.Srctype))
	}

	for i := len(srcTypes); i > 1; i-- {
		subtargets = append(subtargets, subtargets[:precnt])
	}

	for i, s := range subtargets {
		s.sourcetype = srcTypes[i / precnt]
	}

	precnt = len(subtargets)

	geo := make([]string)
	if len(camp.Target.City) > 0 || len(camp.Target.Region > 0 {
		geo = append(geo, camp.Target.City)
		geo = append(geo, camp.Target.Region)
	} else if len(camp.Target.Country) > 0 {
		geo = append(geo, camp.Target.Country)
	}

	for i := len(geo); i > 1; i-- {
		subtargets = append(subtargets, subtargets[:precnt])
	}

	for i, s := range subtargets {
		s.geo = geo[i / precnt]
	}

	campTarKey := fmt.Sprintf(REDISKEY_CAMPAIGN_TARGET, camp.Id)
	for i, s := range subtargets {
		s.adtype = int32(camp.Basic.AdType)

		cli.SAdd(s.Key(), camp.Id)
		cli.SAdd(campTarKey, s.Key())
	}

	return nil
}


/*
*
*/
func UpdateTarget(cli *redis.Client, campId uint64, target *typedef.Target) error {
	if cli == nil {
		return ErrRedisCliNullPtr
	}
	if target == nil {
		return ErrTargetNullPtr
	}

	{
		key := fmt.Sprintf(REDISKEY_CAMPAIGN_TARGET, campId)
		tarkeys , err := cli.SMembers(key).Result()
		if err != nil {
			return err
		}

		for _ , k := range tarkeys {
			cli.SRem(k, campId)
		}
	}

	{
		key := fmt.Sprintf(REDISKEY_CAMPAIGN_INFO, campId)
		serialize, err := cli.Get(key).Result()
		if err != nil {
			return err
		}

		var newCamp typedef.Campaign
		if err := pb.UnmarshalText(serialize, &newCamp); err != nil {
			return err
		}

		newCamp.Target = target
		cli.Set(key, newCamp.String(), 0)

		return CreateTarget(cli, newCamp)
	}
}


/*
*
*/
func GetTargetCamp(cli *redis.Client,
									adx uint64,
									os  string,
									osv string,
									city		string,
									region	string,
									country string,
									adtype		typedef.AdType,
									devtype		typedef.DevType,
									conntype	typedef.ConnType,
									srctype		typedef.SourceType_eSourceAll) ([]uint64, error) {
	if cli == nil {
		return nil, ErrRedisCliNullPtr
	}

	geocnt := 3
	osvcnt := 2

	subtargets := make([]SubTarget, 0, geocnt * osvcnt)
	geos := []string{city, region, country}
	for i , _ := range geos {
		s := SubTarget {
			adx: adx,
			geo: geos[i],
			adtype:		int32(adtype),
			devtype:  int32(devtype),
			conntype: int32(conntype),
			sourcetype: int32(srctype),
		}

		subtargets = append(subtargets, s)
	}

	cnt := len(subtargets)
	osvv := fmt.Sprintf("%s %s", os, osv)
	osvs := []string{os, osvv}
	for i := len(osvs); i > 1; i-- {
		subtargets = append(subtargets, subtargets[:cnt])
	}

	for i , _ := range subtargets {
		subtargets[i].osv = osvs[i / cnt]
	}

	totalcnt := 0
	resset := make([][]uint64, 0, len(subtargets))
	for _ , s := range subtargets {
		var res []uint64
		if err := cli.SMembers(s.Key()).ScanSlice(res); err == nil {
			resset = append(resset, res)
			totalcnt = totalcnt + len(res)
		}
	}

	if len(resset) <= 0 || totalcnt <= 0 {
		return nil, nil
	}

	{
		res := make([]uint64, 0, totalcnt)

		for i := 1; i < len(resset); i++ {
			for j, _ := range resset[i - 1] {
				for k, _ := range resset[i] {
					if resset[i - 1][j] != resset[i][k] {
						res = append(res, resset[i][k])
					}
				}
			}
		}

		return res, nil
	}
}
