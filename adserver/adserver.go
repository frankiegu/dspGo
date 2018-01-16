package adserver

import (
	_"mdsp/adserver/adx"
	_"github.com/mxmCherry/openrtb"
)



type AdxHandler interface {

}

type AdServer struct {

}

type AdCandidate struct{
	campId uint64
	adType uint64
	creative uint64
}

type Creative struct {
	BannerCreative int
	PopupCreative int
	NativeCreative int
}

type AdCandidates struct{
	AdCandidate []AdCandidate
	Creative Creative
}

/*message AdCandidates {
	message AdCandidate {
		uint64 campId  = 1;
		adtype
		creative
	}

	repeated AdCandidate ads = 1;

	message Creative {
		oneof crv {
			BannerCreative  bannerCrv = 1;
			PopupCreative   popupCrv  = 2;
			NativeCreative  nativeCrv = 3;
			}
			}
			}*/


//would use different transport 




/*
* get candidate campaign by filtering targetting and imp
  fitler by other condition such as ip, black/white list category etc
  calc bid price
  rank candicate campaign
  pick creative if top candidate campaign has multi creative
*
*
*/
/*func(s *AdServer) HandleBidding(_ context.Context, req *openrtb.BidRequest, adx uint64) (adRes AdResults, err error) {
	for k, r := range(req.Imp) {
		if req.Imp[k].Banner != nil {
			//ads, err := s.Client.RetrieveBanner(req, adx, false)
		} else if req.Imp[k].Native != nil {
			//ads, err := s.Client.RetrieveNative(req, adx, false)
		}


	}
}*/
