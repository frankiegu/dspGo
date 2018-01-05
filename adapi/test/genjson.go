package main

import (
	_ "bytes"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"reflect"

	"mdsp/common/typedef"
)

func test_proto() {
	var pc typedef.PopupCreative
	pc.Id = 91
	pc.CamId = 92
	pc.IsActive = true
	pc.Html = "what the fuck"

	var cpc typedef.Creative_PopupCrv
	cpc.PopupCrv = &pc

	var c typedef.Creative
	c.Crv = &cpc

	var arr []*typedef.Creative
	arr = append(arr, &c)

	basic := typedef.CampBasic{}

	var cam typedef.Campaign
	cam.Id = 81
	cam.Basic = &basic
	cam.Creatives = arr

	b, err := json.Marshal(cam)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(b))
	}

	var cam2 typedef.Campaign

	s2 := `{"basic":{"id":81},"creatives":[{"popupCrv":{"id":91,"camId":92,"isActive":true,"html":"what the fuck"}}]}`
	err = jsonpb.UnmarshalString(s2, &cam2)
	if err != nil {
		fmt.Println("err", err)
	} else {
		fmt.Println(reflect.TypeOf(cam2.Creatives[0].Crv))
	}

	fmt.Println(cam.String())

	var m jsonpb.Marshaler
	s3, err := m.MarshalToString(&cam)
	fmt.Println(err, s3)
}

func test6() {
	var a []int
	a = append(a, 9)
	if a != nil {
		fmt.Println("---", len(a))
	}
	s := "[]"
	err := json.Unmarshal([]byte(s), &a)
	if err != nil {
		fmt.Println("err", err)
	} else {
		if a == nil {
			fmt.Println("nil")
		} else {
			fmt.Println("not nil", len(a))
		}
	}
}

func get_basic() *typedef.CampBasic {
	var c typedef.CampBasic
	c.Name = "camp_name"
	c.AdType = typedef.AdType_eNative
	c.IsActive = true
	c.Trkimpurl = "trkimpurl"
	c.Trkcampurl = "trkcampurl"
	c.Advtdomain = "advtdomain.com"
	c.Payout = 1000000
	c.PayoutMode = typedef.PayoutMode_eCpaDynamic
	c.Convurl = "convurl"
	c.MaxBidPrice = 9000000
	c.Userid = 123
	return &c
}

func get_budget() *typedef.Budget {
	var b typedef.Budget
	b.UnlimitedEnable = true
	b.TotalBudget = 10000000
	b.DailyBudget = 2000000
	b.PlacementBudget = 1000000
	b.StartStamp = 1505250738
	b.EndStamp = 1505270738
	b.FreqCappingEnable = true
	b.FreqCapping = 1
	b.FreqCappingInterval = typedef.FreqCappingInterval_e24Hour
	b.SpendModel = typedef.SpendModel_eSmooth
	b.DayParting = []uint64{16777215, 16777215, 16777215, 16777214, 16777213, 16777212, 15777214}
	b.Timezone = "Asia/Shanghai"
	return &b
}

func getTarget() *typedef.Target {
	var t typedef.Target
	t.Adxs = []uint64{3, 2}
	t.Categories = []string{"tcat1", "tcat2"}
	t.Country = []string{"French", "Italy", "German"}
	t.Region = []string{"Region1", "Region2"}
	t.City = []string{"city1", "city2"}
	t.Devtype = []typedef.DevType{typedef.DevType_eDevTablet, typedef.DevType_eDevMobile, typedef.DevType_eDevDesktop}
	t.Conntype = typedef.ConnType_eConnWifi
	t.Carrier = []string{"carrier1", "carrier2"}
	t.Osv = []string{"osversion1", "osversion2"}
	t.Ips = []string{"127.0.0.1"}
	t.Autype = typedef.AudienceType_eAudienceBlack
	t.RetargetingAuListId = []string{"retarau1", "retarau2"}
	t.IsRetargettingEnable = true
	t.ViewerListName = "viewName"
	t.VisitorListName = "visitName"
	t.ConverterListName = "convertName"
	t.InvenType = typedef.InventoryType_eInventoryBlack
	t.InvenName = []string{"inven1", "inven2"}
	t.IsIdfaGaidValid = true
	t.Srctype = typedef.SourceType_eSourceInApp
	return &t
}

func getBannerCreative(id uint64) *typedef.BannerCreative {
	var b typedef.BannerCreative
	b.Id = id
	b.CamId = 1
	b.IsActive = true
	b.Mime = typedef.ImageMime_eJPG
	b.Imgurl = "imgurl"
	b.Width = 20
	b.Height = 10
	// snippet
	b.Snippet = getSnippet()
	return &b
}

func getSnippet() *typedef.CreativeSnippet {
	var c typedef.CreativeSnippet
	c.Adm = "adm"
	c.Adomain = []string{"domain1", "domain2"}
	c.Nurl = "nurl"
	c.Iurl = "iurl"
	c.FlowId = 61
	c.Desturl = "desturl"
	c.Campurl = "campurl"
	return &c
}

func getPopCreative(id uint64) *typedef.PopupCreative {
	var p typedef.PopupCreative
	p.Id = id
	p.CamId = 2
	p.IsActive = true
	p.Html = "html"
	p.Snippet = getSnippet()
	return &p
}

func getNative() *typedef.Native {
	var asset typedef.Native_Asset

	var title typedef.Native_Title
	title.Len = 108
	title.Text = "txt"
	asset.Title = append(asset.Title, &title)

	var data typedef.Native_Data
	data.Type = 2
	data.Len = 208
	data.Value = "value"
	asset.Data = append(asset.Data, &data)

	var image typedef.Native_Image
	image.W = 33
	image.H = 44
	image.Mime = typedef.ImageMime_eJPG
	image.Url = "url"
	image.Type = typedef.Native_eNativeImageMain
	asset.Image = append(asset.Image, &image)

	var v typedef.Native_Video
	v.W = 41
	v.H = 42
	v.Duration = 43
	v.Mime = "video/mpg"
	v.Url = "url"
	v.CoverUrl = "coverurl"
	asset.Video = append(asset.Video, &v)

	var l typedef.Native_Link
	l.Url = "url"
	l.Fallback = "fallback"
	l.Clicktrackers = []string{"clicktrackers"}

	var n typedef.Native
	n.Asset = &asset
	n.Link = &l

	return &n
}

func getNativeCreative(id uint64) *typedef.NativeCreative {
	var nc typedef.NativeCreative
	nc.Id = 51
	nc.CamId = 3
	nc.IsActive = true
	nc.Native = getNative()
	nc.Snippet = getSnippet()

	return &nc
}

func getCreativeArray() (arr []*typedef.Creative) {
	var c1, c2, c3 typedef.Creative
	var d1 typedef.Creative_BannerCrv
	var (
		d2 typedef.Creative_PopupCrv
		d3 typedef.Creative_NativeCrv
	)

	d1.BannerCrv = getBannerCreative(2)
	d2.PopupCrv = getPopCreative(3)
	d3.NativeCrv = getNativeCreative(4)
	c1.Crv = &d1
	c2.Crv = &d2
	c3.Crv = &d3

	arr = append(arr, &c1, &c2, &c3)
	return
}

func getCamp() *typedef.Campaign {
	var c typedef.Campaign
	c.Id = 7
	c.Basic = get_basic()
	c.Budget = get_budget()
	c.Target = getTarget()
	c.Creatives = getCreativeArray()

	return &c
}

func testMarshal() {
	c := getCamp()
	var m jsonpb.Marshaler
	s, err := m.MarshalToString(c)
	if err != nil {
		fmt.Println("err", err)
	} else {
		fmt.Println(s)
	}
}

func main() {
	testMarshal()
}
