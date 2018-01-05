package view

import (
	//"encoding/json"
	"fmt"
	//"github.com/golang/protobuf/jsonpb"
	//"github.com/golang/protobuf/proto"
	//"io/ioutil"
	"log"
	//"net/http"

	//"mdsp/common/consts"
	"mdsp/common/typedef"
	//"mdsp/utils/rabbitmq"
)

const (
	assignBudgetTemp         = "[assignBudget]%v"
	assignTargetTemp         = "[assignTarget]%v"
	assignCampaignTemp       = "[assignCampaign]%v"
	assignCampBasicTemp      = "[assignCampBasic]%v"
	assignSnippetTemp        = "[assignSnippet]%v"
	assignBannerCreativeTemp = "[assignBannerCreative]%v"
	assignPopupCreativeTemp  = "[assignPopupCreative]%v"
	assignNativeCreativeTemp = "[assignNativeCreative]%v"
	assignNativeTemp         = "[assignNative]%v"
	assignNativeAssetTemp    = "[assignNativeAsset]%v"
	assignNativeTitleTemp    = "[assignNativeTitle]%v"
	assignNativeImageTemp    = "[assignNativeImage]%v"
	assignNativeDataTemp     = "[assignNativeData]%v"
	assignNativeVideoTemp    = "[assignNativeVideo]%v"
	assignNativeLinkTemp     = "[assignNativeLink]%v"
	appendBannerCreativeTemp = "[appendBannerCreative]%v"
	appendPopupCreativeTemp  = "[appendPopupCreative]%v"
	appendNativeCreativeTemp = "[appendNativeCreative]%v"
)

// assign from jbPtr to pbPtr
func assignBudget(jbPtr *Budget, pbPtr *typedef.Budget) (err error) {
	if jbPtr == nil || pbPtr == nil {
		err = fmt.Errorf(assignBudgetTemp, "jsonBudgetPtr or protoBudgetPtr can't be nil")
		return
	}

	pbPtr.UnlimitedEnable = jbPtr.UnlimitedEnable
	pbPtr.TotalBudget = jbPtr.TotalBudget
	pbPtr.DailyBudget = jbPtr.DailyBudget
	pbPtr.PlacementBudget = jbPtr.PlacementBudget
	pbPtr.StartStamp = jbPtr.StartStamp
	pbPtr.EndStamp = jbPtr.EndStamp
	pbPtr.FreqCappingEnable = jbPtr.FreqCappingEnable
	pbPtr.FreqCapping = jbPtr.FreqCap

	// enum: freq cap interval: from string to int32
	pbPtr.FreqCappingInterval = 1 // 3hour

	// enum: spend model: from string to int32
	pbPtr.SpendModel = 1 // asap

	pbPtr.DayParting = jbPtr.DayParting
	pbPtr.Timezone = jbPtr.Timezone

	return
}

func assignTarget(jtPtr *Target, ptPtr *typedef.Target) (err error) {
	if jtPtr == nil || ptPtr == nil {
		err = fmt.Errorf(assignTargetTemp, "jsonTargetPtr or protoTargetPtr can't be nil")
		return
	}

	ptPtr.Adxs = jtPtr.Adxs
	ptPtr.Categories = jtPtr.Categories
	ptPtr.Country = jtPtr.Country
	ptPtr.Region = jtPtr.Region
	ptPtr.City = jtPtr.City

	// enum array
	// ptPtr.Devtype = jtPtr.Devtype // convert json string to pb enu

	// enum
	ptPtr.Conntype = 1 // 0:unknown, 1:all, 2:mobile, 3:wifi

	ptPtr.Carrier = jtPtr.Carrier
	ptPtr.Osv = jtPtr.Osv
	ptPtr.Ips = jtPtr.Ips

	// enum
	ptPtr.Autype = 1 // 0:nontype, 1:white, 2:black
	ptPtr.RetargetingAuListId = jtPtr.RetargetingAuListId

	ptPtr.ViewerListName = jtPtr.ViewerListName
	ptPtr.VisitorListName = jtPtr.VisitorListName
	ptPtr.ConverterListName = jtPtr.ConverterListName
	if jtPtr.ViewerListName != "" || jtPtr.VisitorListName != "" || jtPtr.ConverterListName != "" {
		ptPtr.IsRetargettingEnable = true
	} else {
		ptPtr.IsRetargettingEnable = false
	}

	// enum
	ptPtr.InvenType = 1 // 0:unknown, 1:white, 2:black
	ptPtr.InvenName = jtPtr.InvenName

	ptPtr.IsIdfaGaidValid = jtPtr.IsIdfaGaidValid

	// enum
	ptPtr.Srctype = 1 // 0:unknown, 1:app, 2:web

	return
}

func assignCampaign(jcamPtr *Campaign, pcamPtr *typedef.Campaign) (err error) {
	if jcamPtr == nil || pcamPtr == nil {
		err = fmt.Errorf(assignCampaignTemp, "jcamPtr or pcamPtr can't be nil")
		return
	}

	// camp id
	if jcamPtr.Id == 0 {
		err = fmt.Errorf(assignCampaignTemp, "campId is missing")
		return
	}
	pcamPtr.Id = jcamPtr.Id

	if jcamPtr.Basic != nil {
		var pBasic typedef.CampBasic
		err = assignCampBasic(jcamPtr.Basic, &pBasic)
		if err != nil {
			log.Printf(assignCampaignTemp, err)
			return
		}
		pcamPtr.Basic = &pBasic
	}

	if jcamPtr.Budget != nil {
		var pBudget typedef.Budget
		err = assignBudget(jcamPtr.Budget, &pBudget)
		if err != nil {
			log.Printf(assignCampaignTemp, err)
			return
		}
		pcamPtr.Budget = &pBudget // budget
	}

	if jcamPtr.Target != nil {
		var pTarget typedef.Target
		err = assignTarget(jcamPtr.Target, &pTarget)
		if err != nil {
			log.Printf(assignCampaignTemp, err)
			return
		}
		pcamPtr.Target = &pTarget // target
	}

	// creatives
	var creatives []*typedef.Creative
	if jcamPtr.BannerCreatives != nil && len(jcamPtr.BannerCreatives) != 0 {
		err = appendBannerCreative(jcamPtr.BannerCreatives, &creatives)
		if err != nil {
			err = fmt.Errorf(assignCampaignTemp, err)
			return
		}
	}
	if jcamPtr.PopupCreatives != nil && len(jcamPtr.PopupCreatives) != 0 {
		err = appendPopupCreative(jcamPtr.PopupCreatives, &creatives)
		if err != nil {
			err = fmt.Errorf(assignCampaignTemp, err)
			return
		}
	}
	if jcamPtr.NativeCreatives != nil && len(jcamPtr.NativeCreatives) != 0 {
		err = appendNativeCreative(jcamPtr.NativeCreatives, &creatives)
		if err != nil {
			err = fmt.Errorf(assignCampaignTemp, err)
			return
		}
	}

	pcamPtr.Creatives = creatives

	return
}

func assignCampBasic(jcamPtr *CampBasic, pcamPtr *typedef.CampBasic) (err error) {
	if jcamPtr == nil || pcamPtr == nil {
		err = fmt.Errorf(assignCampBasicTemp, "jcs or pcs can't be nil")
		return
	}
	pcamPtr.Name = jcamPtr.Name

	// enum
	pcamPtr.AdType = 3 // 0:unknown, 1:banner, 2:popup, 3:native

	pcamPtr.IsActive = jcamPtr.IsActive
	pcamPtr.Trkimpurl = jcamPtr.Trkimpurl
	pcamPtr.Trkcampurl = jcamPtr.Trkcampurl
	pcamPtr.Advtdomain = jcamPtr.Advtdomain
	pcamPtr.Payout = jcamPtr.Payout

	// enum
	pcamPtr.PayoutMode = 2 // 0:unknown, 1:fixed, 2:dynamic

	pcamPtr.Convurl = jcamPtr.Convurl
	pcamPtr.MaxBidPrice = jcamPtr.MaxBidPrice
	pcamPtr.Userid = jcamPtr.Userid

	return
}

func assignBannerCreative(jBCrt *BannerCreative, pBCrt *typedef.BannerCreative) (err error) {
	if jBCrt == nil || pBCrt == nil {
		err = fmt.Errorf(assignBannerCreativeTemp, "jBCrt or pBCrt can't be nil")
		return
	}

	// enum
	pBCrt.Mime = 1 // 0:unknown, 1:jpg, 2:png, 3gif

	pBCrt.Imgurl = jBCrt.Imgurl
	pBCrt.Width = jBCrt.Width
	pBCrt.Height = jBCrt.Height

	pBCrt.Id = jBCrt.Id
	pBCrt.CamId = jBCrt.CamId
	pBCrt.IsActive = jBCrt.IsActive

	if jBCrt.Snippet != nil {
		var pSnippet typedef.CreativeSnippet
		err = assignSnippet(jBCrt.Snippet, &pSnippet)
		if err != nil {
			err = fmt.Errorf(assignBannerCreativeTemp, err)
			return
		}
		pBCrt.Snippet = &pSnippet
	}

	return
}

func assignPopupCreative(jPCrt *PopupCreative, pPCrt *typedef.PopupCreative) (err error) {
	if jPCrt == nil || pPCrt == nil {
		err = fmt.Errorf(assignPopupCreativeTemp, "jPCrt or pPCrt can't be nil")
		return
	}

	pPCrt.Id = jPCrt.Id
	pPCrt.CamId = jPCrt.CamId
	pPCrt.IsActive = jPCrt.IsActive
	pPCrt.Html = jPCrt.Html

	if jPCrt.Snippet != nil {
		var pSnippet typedef.CreativeSnippet
		err = assignSnippet(jPCrt.Snippet, &pSnippet)
		if err != nil {
			err = fmt.Errorf(assignPopupCreativeTemp, err)
			return
		}
		pPCrt.Snippet = &pSnippet
	}

	return
}

func assignNativeCreative(jNCrt *NativeCreative, pNCrt *typedef.NativeCreative) (err error) {
	if jNCrt == nil || pNCrt == nil {
		err = fmt.Errorf(assignNativeCreativeTemp, "jNCrt or pNCrt can't be nil")
		return
	}

	pNCrt.Id = jNCrt.Id
	pNCrt.CamId = jNCrt.CamId
	pNCrt.IsActive = jNCrt.IsActive

	if jNCrt.Native != nil {
		var pn typedef.Native
		err = assignNative(jNCrt.Native, &pn)
		if err != nil {
			err = fmt.Errorf(assignNativeCreativeTemp, err)
			return
		}
		pNCrt.Native = &pn
	}

	if jNCrt.Snippet != nil {
		var ps typedef.CreativeSnippet
		err = assignSnippet(jNCrt.Snippet, &ps)
		if err != nil {
			err = fmt.Errorf(assignNativeCreativeTemp, err)
			return
		}
		pNCrt.Snippet = &ps
	}

	return
}

func assignSnippet(jcs *CreativeSnippet, pcs *typedef.CreativeSnippet) (err error) {
	if jcs == nil || pcs == nil {
		err = fmt.Errorf(assignSnippetTemp, "jcs or pcs can't be nil")
		return
	}

	pcs.Adm = jcs.Adm
	pcs.Adomain = jcs.Adomain
	pcs.Nurl = jcs.Nurl
	pcs.Iurl = jcs.Iurl
	pcs.FlowId = jcs.FlowId
	pcs.Desturl = jcs.Desturl
	pcs.Campurl = jcs.Campurl

	return
}

func assignNative(jn *Native, pn *typedef.Native) (err error) {
	if jn == nil {
		return
	}

	if pn == nil {
		err = fmt.Errorf(assignNativeTemp, "pn can't be nil")
		return
	}

	// assign Native_Asset
	if jn.Asset != nil {
		var asset typedef.Native_Asset
		err = assignNativeAsset(jn.Asset, &asset)
		if err != nil {
			err = fmt.Errorf(assignNativeTemp, err)
			return
		}
	}

	// assign Native_Link
	if jn.Link != nil {
		var link typedef.Native_Link
		err = assignNativeLink(jn.Link, &link)
		if err != nil {
			err = fmt.Errorf(assignNativeTemp, err)
			return
		}
	}

	return
}

func assignNativeAsset(jna *Native_Asset, pna *typedef.Native_Asset) (err error) {
	if jna == nil {
		return
	}

	if pna == nil {
		err = fmt.Errorf(assignNativeAssetTemp, "pna can't be nil")
		return
	}

	if jna.Title != nil && len(jna.Title) != 0 {
		var titles []*typedef.Native_Title

		for _, jt := range jna.Title {
			var nt typedef.Native_Title
			err = assignNativeTitle(jt, &nt)
			if err != nil {
				err = fmt.Errorf(assignNativeAssetTemp, err)
				return
			}

			titles = append(titles, &nt)
		}
		pna.Title = titles
	}

	if jna.Data != nil && len(jna.Data) != 0 {
		var datas []*typedef.Native_Data

		for _, jd := range jna.Data {
			var nd typedef.Native_Data
			err = assignNativeData(jd, &nd)
			if err != nil {
				err = fmt.Errorf(assignNativeAssetTemp, err)
				return
			}
			datas = append(datas, &nd)
		}
		pna.Data = datas
	}

	if jna.Image != nil && len(jna.Image) != 0 {
		var images []*typedef.Native_Image

		for _, ji := range jna.Image {
			var ni typedef.Native_Image
			err = assignNativeImage(ji, &ni)
			if err != nil {
				err = fmt.Errorf(assignNativeAssetTemp, err)
				return
			}
			images = append(images, &ni)
		}
		pna.Image = images
	}

	if jna.Video != nil && len(jna.Video) != 0 {
		var videos []*typedef.Native_Video
		for _, jv := range jna.Video {
			var nv typedef.Native_Video
			err = assignNativeVideo(jv, &nv)
			if err != nil {
				err = fmt.Errorf(assignNativeAssetTemp, err)
				return
			}
			videos = append(videos, &nv)
		}
		pna.Video = videos
	}

	return
}

func assignNativeTitle(jnt *Native_Title, pnt *typedef.Native_Title) (err error) {
	if jnt == nil {
		return
	}

	if pnt == nil {
		err = fmt.Errorf(assignNativeTitleTemp, "pnt can't be nil")
		return
	}

	pnt.Len = jnt.Len
	pnt.Text = jnt.Text

	return
}

func assignNativeImage(jni *Native_Image, pni *typedef.Native_Image) (err error) {
	if jni == nil {
		return
	}

	if pni == nil {
		err = fmt.Errorf(assignNativeImageTemp, "pni can't be nil")
		return
	}

	pni.W = jni.W
	pni.H = jni.H

	// enum
	pni.Mime = 1 // 0:unknown, 1:jpg, 2:png, 3:gif

	pni.Url = jni.Url

	// enum
	pni.Type = 3 // 0:unknown, 1:icon, 2:logo, 3:main

	return
}

func assignNativeData(jnd *Native_Data, pnd *typedef.Native_Data) (err error) {
	if jnd == nil {
		return
	}

	if pnd == nil {
		err = fmt.Errorf(assignNativeDataTemp, "pnd can't be nil")
		return
	}

	pnd.Type = jnd.Type
	pnd.Len = jnd.Len
	pnd.Value = jnd.Value

	return
}

func assignNativeVideo(jnv *Native_Video, pnv *typedef.Native_Video) (err error) {
	if jnv == nil {
		return
	}

	if pnv == nil {
		err = fmt.Errorf(assignNativeVideoTemp, "pnv can't be nil")
		return
	}

	pnv.W = jnv.W
	pnv.H = jnv.H
	pnv.Duration = jnv.Duration
	pnv.Mime = jnv.Mime
	pnv.Url = jnv.Url
	pnv.CoverUrl = jnv.CoverUrl

	return
}

func assignNativeLink(jnl *Native_Link, pnl *typedef.Native_Link) (err error) {
	if jnl == nil {
		return
	}

	if pnl == nil {
		err = fmt.Errorf(assignNativeLinkTemp, "pnl can't be nil")
		return
	}

	pnl.Url = jnl.Url
	pnl.Fallback = jnl.Fallback
	pnl.Clicktrackers = jnl.Clicktrackers

	return
}

func appendBannerCreative(jc []*BannerCreative, pcPtr *[]*typedef.Creative) (err error) {
	if jc == nil || len(jc) == 0 {
		return
	}

	if pcPtr == nil {
		err = fmt.Errorf(appendBannerCreativeTemp, "pcPtr can't be nil")
		return
	}

	for _, j := range jc {
		var p typedef.BannerCreative
		err = assignBannerCreative(j, &p)
		if err != nil {
			err = fmt.Errorf(appendBannerCreativeTemp, err)
			return
		}

		var bc typedef.Creative_BannerCrv
		bc.BannerCrv = &p

		var c typedef.Creative
		c.Crv = &bc
		*pcPtr = append(*pcPtr, &c)
	}

	return
}

func appendPopupCreative(jc []*PopupCreative, pcPtr *[]*typedef.Creative) (err error) {
	if jc == nil || len(jc) == 0 {
		return
	}

	if pcPtr == nil {
		err = fmt.Errorf(appendPopupCreativeTemp, "pcPtr can't be nil")
		return
	}

	for _, j := range jc {
		var p typedef.PopupCreative
		err = assignPopupCreative(j, &p)
		if err != nil {
			err = fmt.Errorf(appendPopupCreativeTemp, err)
			return
		}

		var cc typedef.Creative_PopupCrv
		cc.PopupCrv = &p

		var c typedef.Creative
		c.Crv = &cc
		*pcPtr = append(*pcPtr, &c)
	}

	return
}

func appendNativeCreative(jc []*NativeCreative, pcPtr *[]*typedef.Creative) (err error) {
	if jc == nil || len(jc) == 0 {
		return
	}

	if pcPtr == nil {
		err = fmt.Errorf(appendNativeCreativeTemp, "pcPtr can't be nil")
		return
	}

	for _, j := range jc {
		var p typedef.NativeCreative
		err = assignNativeCreative(j, &p)
		if err != nil {
			err = fmt.Errorf(appendNativeCreativeTemp, err)
			return
		}

		var cc typedef.Creative_NativeCrv
		cc.NativeCrv = &p

		var c typedef.Creative
		c.Crv = &cc
		*pcPtr = append(*pcPtr, &c)
	}

	return
}
