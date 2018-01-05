package view

type Budget struct {
	UnlimitedEnable bool   `json:"unlimitedBudget"`
	TotalBudget     uint64 `json:"totalBudget"`
	DailyBudget     uint64 `json:"dailyBudget"`
	PlacementBudget uint64 `json:"dailyPerPlacementBudget"`
	SpendModel      string `json:"spendStrategy"` // SMOOTH, ASAP

	FreqCappingEnable   bool   `json:"freqCapEnabled"`
	FreqCap             uint64 `json:"freqCountLimit"`
	FreqCappingInterval string `json:"freqTimeWindow"` // 24hours, 6hours

	StartStamp uint64 `json:"activityPeriodsStartTime"`
	EndStamp   uint64 `json:"activityPeriodsEndTime"`

	// From Mon to Sun
	DayParting []uint64 `json:"dayParting"`
	Timezone   string   `json:"activityPeriodsTimezone"`
}

type Target struct {
	Adxs []uint64 `json:"adExchanges"` // adx_id list
	// 2 level category
	Categories []string `json:"categories"` // IAB22-4 list

	// geo
	Country []string `json:"countries"`
	Region  []string `json:"states"`
	City    []string `json:"cities"`

	// device type
	Devtype []string `json:"deviceTypes"` // mobile, tablet, desktop

	// connection type
	Conntype string `json:"connectionType"` // WIFI, ANY, MOBILE

	Carrier []string `json:"carriers"` // AWCC, Roshan
	// os version
	Osv []string `json:"oses"`

	Ips []string `json:"ips"`

	// audience type
	Autype              string   `json:"audienceType"` // BLACK_LIST, WHITE_LIST
	RetargetingAuListId []string `json:"audienceIds"`

	//IsRecordEnable         bool   `json:""`
	ViewerListName    string `json:"targetingAudienceViewer"` // 0: not assigned
	VisitorListName   string `json:"targetingAudienceVistor"`
	ConverterListName string `json:"targetingAudienceConverter"`

	InvenType string   `json:"clientType"` // BLACK_LIST, WHITE_LIST
	InvenName []string `json:"clientIds"`

	IsIdfaGaidValid bool   `json:"isIdfaGaid"`
	SrcType         string `json:"sourceType"` // ANY, SITE, APP
}

// front end check
type Creative struct {
	Id       uint64 `json:"id"`
	CamId    uint64 `json:"campId"` // check
	IsActive bool   `json:"status"`
}

type CreativeSnippet struct { // front need to add
	Adm     string   `json:"adm"`
	Adomain []string `json:"adomain"`
	Nurl    string   `json:"nurl"`
	Iurl    string   `json:"iurl"`
	FlowId  uint64   `json:"flowId"`
	Desturl string   `json:"destUrl"`
	Campurl string   `json:"campUrl"`
}

type BannerCreative struct {
	Creative
	Snippet *CreativeSnippet `json:"snippet"`

	Mime   string `json:"mime"`
	Imgurl string `json:"imgurl"`
	Width  uint32 `json:"width"`
	Height uint32 `json:"height"`
}

type PopupCreative struct {
	Creative
	Snippet *CreativeSnippet `json:"snippet"`

	Html string `json:"html"`
}

type Native_Title struct {
	Len  uint64 `json:"len"`
	Text string `json:"text"`
}

type Native_Data struct {
	Type  int64  `json:"type"`
	Len   uint64 `json:"len"`
	Value string `json:"value"`
}

type Native_Image struct {
	W    int32  `json:"width"`
	H    int32  `json:"height"`
	Mime string `json:"mime"`
	Url  string `json:"url"`
	Type string `json:"type"`
}

type Native_Video struct {
	W        int32  `json:"width"`
	H        int32  `json:"height"`
	Duration int32  `json:"duration"`
	Mime     string `json:"mime"`
	Url      string `json:"url"`
	CoverUrl string `json:"coverUrl"`
}

type Native_Asset struct {
	Title []*Native_Title `json:"title"`
	Data  []*Native_Data  `json:"data"`
	Image []*Native_Image `json:"image"`
	Video []*Native_Video `json:"video"`
}

type Native_Link struct {
	Url           string   `json:"url"`
	Fallback      string   `json:"fallback"`
	Clicktrackers []string `json:"clicktrackers"`
}

type Native struct {
	Asset *Native_Asset `json:"asset"`
	Link  *Native_Link  `json:"link"`
}

type NativeCreative struct {
	Creative
	Native  *Native          `json:"native"`
	Snippet *CreativeSnippet `json:"snippet"`
}

type CampBasic struct {
	Name string `json:"name"`

	AdType   string `json:"type"`
	IsActive bool   `json:"status"` // 0:paused,1:active

	Trkimpurl  string `json:"trkImpUrl"`
	Trkcampurl string `json:"trkCampUrl"`
	Advtdomain string `json:"domain"`

	PayoutMode string `json:"revenueType"` // CPAFIXED,CPADYNAMIC
	// payout 10e6 in uint
	Payout uint64 `json:"revenueValue"`
	// conversion action in url
	Convurl string `json:"conversionActionUrl"`
	// max bid price 10e6 in uint
	MaxBidPrice uint64 `json:"bidPrice"`
	// user id of campaign
	Userid uint64 `json:"userId"`
}

type Campaign struct {
	Id     uint64     `json:"campId"`
	Basic  *CampBasic `json:"basic"`
	Budget *Budget    `json:"budget"`
	Target *Target    `json:"target"`

	// creatives
	BannerCreatives []*BannerCreative `json:"banners"`
	PopupCreatives  []*PopupCreative  `json:"popups"`
	NativeCreatives []*NativeCreative `json:"natives"`
}

/*
type BudgetUpdate struct {
	Id     uint64  `json:"campId"`
	Budget *Budget `json:"budget"`
}
*/
