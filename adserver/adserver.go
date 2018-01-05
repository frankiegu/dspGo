package adserver


import (

	"dsp/adserver/adx"
)

type AdServer struct {

}


type AdxHandler interface {
}


message AdCandidates {
	message AdCandidate {
		uint64 campId  = 1;
		adtype
		creative
	}

	repeated AdCandidate	ads = 1;

	message Creative {
  oneof crv {
    BannerCreative  bannerCrv = 1;
    PopupCreative   popupCrv  = 2;
    NativeCreative  nativeCrv = 3;
  }
}


}

//would use different transport 
type AdRetriever interface {
	RetrieveBanner(cli, req *openrtb.Request) (AdCandidates, error)
	RetrievePopup(cli, req *openrtb.Request) (AdCandidates,  error)
	RetrieveNative(cli, req *openrtb.Request) (AdCandidates, error)
}



/*
* get candidate campaign by filtering targetting and imp
  fitler by other condition such as ip, black/white list category etc
  calc bid price
  rank candicate campaign
  pick creative if top candidate campaign has multi creative
*
*
*/
func(s *AdServer) HandleBidding(_ context.Context, req *openrtb.BidRequest, adx uint64) (adRes AdResults, err error) {
	for , r := req.Imp {
		if req.Imp.Banner != nil {
			ads, err := s.Client.RetrieveBanner(req, adx, false)
		} else if req.Imp.Native != nil {
			ads, err := s.Client.RetrieveNative(req, adx, false)
		}


	}
}
