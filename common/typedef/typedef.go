package typedef

/*
* api2be message package
* format: json
* eg:
* 	{
*			key  : "startcampaign",
*     body : 111111
*   }
 */

type AdApi2BeMsg struct {
	Key		string
	Body	string
}

/*
*	event2cc message package
* format: json
* eg:
*
 */
type AdEvent2ccMsg struct {
}

/*
* cc2bidder messsage package
* Key:
*				CC_KEY_ACTIVECAMPAIGN
*				CC_KEY_INACTIVECAMPAIGN
*
* CampIds: array of campaign id
 */

type AdCC2BidderMsg struct {
	Key     string
	CampIds []uint64
}

//api2be
type CampMsg struct {
	Id   uint64 `json:"id"`
	Body string `json:"body"`
}

type CampIdMsg struct {
	Id		uint64 `json:"id"`
}

/*
type BudgetMsg struct {
	CampId uint64 `json:"campId"`
	Budget *Budget `json:"budget"`
}
*/

type CreativeMsg struct {
	CampId uint64 `json:"campId"`
	//
	CreativeIds []uint64 `json:"creativeIds"`
}

type InvenMsg struct {
	CampId uint64 `json:"campId"`
	// InvenType
	InvenNames []string `json:"invenNames"`
}

type UpdateRetargeting struct {
	CampId   uint64   `json:"campId"`
	AuType   string   `json:"auType"`
	AuIdList []string `json:"auIdList"`
}

type UpdateCreative struct {
	Crves	[]string	`json:"crves"`
}


