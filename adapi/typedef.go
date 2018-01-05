package adapi

type Campaign struct {
	Id			uint64		`json:"campaignId"`
	Adtype	string    `json:"type"`  //NATIVE, BANNER, POPUP

	Basic		*CampBasic	`json:"basic"`
	Target	*CampTarget	`json:"target"`
	Budget	*CampBudget	`json:"budget"`

	Popups  []PopupCrv	`json:"popups"`
	Banners []BannerCrv `json:"banners"`
	Natives []NativeCrv	`json:"natives"`
}



type CampBasic struct {
	Name	string		`json:"name"`

	IsActive	int		 `json:"status"`
	AdvDomain	string `json:"domain"`

	MaxBidPrice	uint64		`json:"bidPrice"`
	UID					uint64		`json:"userId"`

	PayoutMode  string    `json:"revenueType"` //CPAFIXED,CPADYNAMIC
	Payout			uint64		`json:"revenueValue"`

	ConvUrl			string		`json:"conversionActionUrl"`

	TrkCampUrl	string		`json:"trackingCampaignUrl"`
	TrkImpUrl		string		`json:"trackingCampaignImpressionUrl"`
	TrkClickUrl	string		`json:"trackingCampaignClickUrl"`
}


type CampTarget struct {
	Adx		[]uint64					`json:"adx"`

	Category	[]string		`json:"categories"`
	Country		[]string		`json:"countries"`
	Region		[]string		`json:"states"`
	City			[]string		`json:"cites"`

	Carrier		[]string		`json:"carriers"`
	OSV				[]string		`json:"oses"`
	IPS				[]string		`json:"ips"`

	ConnType		string				`json:"connectionType"`	//WIFI,ANY,MOBILE
	DevType			[]string			`json:"deviceTypes"`	  //TABLET,MOBILE,DESKTOP
	SourceType	[]string			`json:"sourceType"`     //ANY,SITE,APP

	AuType	   string			`json:"audienceType"`  //BLACK_LIST,WHITE_LIST
	AuList	   []string		`json:"audienceIds"`

	InvenType		string			`json:"clientType"` //BLACK_LIST,WHITE_LIST
	InvenList		[]string	  `json:"clientIds"`

	IsRetargettingEnable	int		`json:"isretargetting"`

	ViewerListName		string		`json:"viewerlist"`
	VisitorListName		string		`json:"visitorlist"`
	ConverterListName	string		`json:"converterlsit"`

	IsIdfaGaidValid	int	`json:"isIdfaGaid"`
}

type CampBudget struct {
	IsUnlimitedEnable	int		`json:"unlimitedBudget"`

	TotalBudget		uint64		`json:"totalBudget"`
	DailyBudget		uint64		`json:"dailyBudget"`
	PlacementBudget	uint64	`json:"dailyPerPlacementBudget"`

	StartTime			string	`json:"activityPeriodsStartTime"`
	EndTime				string	`json:"activityPeriodsEndTime"`

	IsFreqCappingEnable	int		`json:"freqCapEnabled"`
	FreqCapping			uint64		`json:"freqCountLimit"`
	FreqInterval		int32			`json:"freqTimeWindow"`

	SpendModel			string		`json:"spendStrategy"`			//SMOOTH , ASAP

	TimeZone				string		`json:"activityPeriodsTimezone"`
	DayParting			[][]int		`json:"dayParting"`

}

type CampRetarget struct {
	IsRetargettingEnable	int		`json:"isretargetting"`

	ViewerListName		string		`json:"viewerlist"`
	VisitorListName		string		`json:"visitorlist"`
	ConverterListName	string		`json:"converterlist"`
}

type PopupCrv struct {
	Id		uint64	`json:"creativeId"`
	Url		string	`json:"url"`

	IsApproved	int		`json:"isApproved"`
	IsActive	int			`json:"status"`
	FlowId		int			`json:"flowId"`
	DestUrl		string	`json:"redirectUrl"`
}

type ImageInfo  struct {
	W			int			`json:"width"`
	H			int			`json:"height"`

	Mime	string		`json:"mime"`
	Url		string		`json:"cdnUrl"`
}

type BannerCrv struct {
	Id		uint64			`json:"creativeId"`

	IsApproved	int		`json:"isApproved"`
	IsActive	int			`json:"status"`
	FlowId		int			`json:"flowId"`
	DestUrl		string	`json:"redirectUrl"`

	Img		*ImageInfo	`json:"image"`
}

type NativeCrv struct {
	Id		uint64	`json:"creativeId"`

  HeadLine	string	`json:"headline"`
	CtaText		string	`json:"ctaText"`

	IconImg *ImageInfo			`json:"iconImage"`
	MainImg *ImageInfo			`json:"mainImage"`

	IsApproved	int		`json:"isApproved"`

	IsActive	int			`json:"status"`
	FlowId		int			`json:"flowId"`
	DestUrl		string	`json:"redirectUrl"`
}

type CampCreateiveId struct {
	Ids		[]uint64	`json:"crvIds"`
}

type Inventory struct {
	InvenType		string			`json:"clientType"` //BLACK_LIST,WHITE_LIST
	InvenList		[]string	  `json:"clientIds"`
}

type Audience struct {
	AuType	   string			`json:"audienceType"`  //BLACK_LIST,WHITE_LIST
	AuList	   []string		`json:"audienceIds"`
}
