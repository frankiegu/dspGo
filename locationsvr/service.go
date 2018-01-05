package locationsvr


import (
	"tracking/utils/ip2location"
	"tracking/locationsvr/proto"

	"tracking/utils/log"
	"tracking/utils/countrycode"
	"golang.org/x/net/context"
)

type LocSvr struct {}

func NewLocSvr() *LocSvr {
		return &LocSvr{}
}

func (s *LocSvr)Ip2Location(ctx context.Context, req *proto.LocRequest) (*proto.IpLocation, error) {
	l := ip2location.Get_all(req.Ip)

	rsp := &proto.IpLocation {
		CountryName: l.Country_long,
		CountryAbbr: countrycode.CountryCode2To3(l.Country_short), //l.Country_short,
		Region:			 l.Region,
		City:				 l.City,

		Isp:				l.Isp,
		Lat:				l.Latitude,
		Lon:				l.Longitude,

		Domain:			l.Domain,
		ZipCode:		l.Zipcode,
		TimeZone:		l.Timezone,
		NetSpeed:		l.Netspeed,
		IddCode:		l.Iddcode,
		AreaCode:		l.Areacode,

		WeatherName:	l.Weatherstationname,
		WeatherCode:	l.Weatherstationcode,

		Mcc:	l.Mcc,
		Mnc:	l.Mnc,

		MobileBrand:	l.Mobilebrand,
	}

	log.Logger().Debugf("location service request ip=%s, resp location=%+v\n", req.Ip, rsp)
	return rsp, nil
}
