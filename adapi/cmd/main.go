/*
* #dsp api service
*
 */

package main

/*
* using http post method to transport campaign package
*
 */

/*
* campaign package format
*
*	type: json
*
	{
		func: "",
		value: {
			param:
		}
	}



*/

/*
* func : CreateCampaign
* param: Campaign
* type : json
* desc : when create campaign 
*
* path : "/api/fe/v1/createcamp"
*/



/*
* func : UpdateCampaignBasic
* param: CampBasic
* type : json
* desc : when update campaign basic 
*
* path : "/api/fe/v1/updatebasic/{id}"
*/


/*
* func : UpdateCampaignTarget
* param: CampTarget
* type : json object
* desc : when update campaign target 
*
* path : "/api/fe/v1/updatetarget/{id}"
*/


/*
* func : UpdateRetargetting
* param: CampRetarget
* desc : when update retargetting
*
* path : "api/fe/v1/updateretargetting/{id}"
*/



/*
* func : UpdateCampaingBudget
* param: CampBudget
* type : json object
* desc : when update campaign budget
*
* path : "/api/fe/v1/udpatebudget/{id}"
*/



/*
* func : UpdateCampaignCreative
* param: creative
* type : json object
* desc : when udpate/create creative
*/

/*
* func : UpdatePopup
* param: []PopupCrv
* desc : 
*				case 1: update oen or more existed popups
*       case 2: create one or more new popups 
*
* path :"/api/fe/v1/updatepopup/{id}"
*/



/*
* func : UpdateBanner
* param: []BannerCrv
* desc : 
*				case 1: update oen or more existed banners
*       case 2: create one or more new banners 
*
* path : "/api/fe/v1/updatebanner/{id}"
*/



/*
* func : UpdateNative
* param: []NativeCrv
* desc : 
*				case 1: update oen or more existed natives
*       case 2: create one or more new natives 
*
* path : "/api/fe/v1/updatenative/{id}"
*/



/*
* func : DeleteCampaign
* param: None
* type : uint64
* desc : campaign id
*
* path : "/api/fe/v1/deletecamp/{id}"
*/

/*
* func : StartCampaign
* param: None
* type : uint64
* desc : start campaign to bid
*
* path : "/api/fe/v1/startcamp/{id}"
*/

/*
* func : PauseCampaign
* param: None
* type : uint64
* desc : pause bidding campaign

* path : "/api/fe/v1/pausecamp/{id}"
*/

/*
* func : ActiveCreative
* param: CampCreateiveId
* type : json object
* desc : 
*				 case 1: inactive one or more creatives of campaign
*
* path : "/api/fe/v1/activecrv/{id}"
*/

// CampCreateiveId

/*
* func : InactiveCreative
* param: CampCreateiveId
* type : json object
* desc : 
*				 case 1: inactive one or more creatives of campaign
*				 case 2: delete one or more creatives of campaign
*
*
* path  : "/api/fe/v1/approvecrv/{id}"
*/

/*
* func : ApproveCreative
* param: CampCreateiveId
* type : json object
* desc : 
*				 case 1: inactive one or more creatives of campaign
*				 case 2: delete one or more creatives of campaign
*
* path : "/api/fe/v1/inactivecrv/{id}"
*/



/*
* func: UpdateInventory
* param : Inventory
* type  : json object
* desc  : update white/black list name
* path  : "/api/fe/v1/updateinventory/{id}"
*/


/*
* func: UpdateAudience
* param : Audience
* type  : json type
* desc  : update include/exclude list name
* path  : "/api/fe/v1/updateaudience/{id}"
*/

import (
	"log"
	//"flag"
	"net/http"

	"mdsp/adapi"
	"github.com/gorilla/mux"
	//"mdsp/utils/conf"
	//"mdsp/common/consts"
	"mdsp/utils/rabbitmq"
)

const (
	GET	 = "GET"
	POST = "POST"
)

var rmq 
//need login or
//delpoy in the same local network
func main() {
	// create campaign
	r := mux.NewRouter()
	r.HandleFunc("/api/fe/v1/createcamp", adapi.CreateCampaignHandler).Methods(POST)

	//r.HandleFunc("/api/fe/v1/{function}/{id}", adapi.FuncV1Handler).Methods(POST)

	// update campaign basic
	r.HandleFunc("/api/fe/v1/updatebasic/{id}", adapi.UpdateBasicHandler).Methods(POST)

	// update campaign targets
	r.HandleFunc("/api/fe/v1/updatetarget/{id}", adapi.UpdateTargetHandler).Methods(POST)

	// update campaign budget 
	r.HandleFunc("api/fe/v1/udpatebudget/{id}", adapi.UpdateBudgetHandler).Methods(POST)

	//popup
	r.HandleFunc("/api/fe/v1/updatepopup/{id}", adapi.UpdatePopupHandler).Methods(POST)

	//banner
	r.HandleFunc("/api/fe/v1/updatebanner/{id}", adapi.UpdateBannerHandler).Methods(POST)

	//native
	r.HandleFunc("/api/fe/v1/updatenative/{id}", adapi.UpdateNativeHandler).Methods(POST)

	// delete campaign
	r.HandleFunc("/api/fe/v1/deletecamp/{id}", adapi.DeleteCampaignHandler).Methods(POST)

	// start campaign
	r.HandleFunc("/api/fe/v1/startcamp/{id}", adapi.StartCampaignHandler).Methods(POST)

	// pause campaign
	r.HandleFunc("/api/fe/v1/pausecamp/{id}", adapi.PauseCampaignHandler).Methods(POST)

	// active creative
	r.HandleFunc("/api/fe/v1/activecrv/{id}", adapi.ActiveCreativeHandler).Methods(POST)

	// inactive creative
	r.HandleFunc("/api/fe/v1/inactivecrv/{id}", adapi.InactiveCreativeHandler).Methods(POST)

	//approve creative
	r.HandleFunc("/api/fe/v1/approvecrv/{id}", adapi.ApproveCreativeHandler).Methods(POST)

	//update inventory
	r.HandleFunc("/api/fe/v1/updateinventory/{id}", adapi.UpdateInventoryHandler).Methods(POST)

	//update audience
	r.HandleFunc("/api/fe/v1/updateaudience/{id}", adapi.UpdateAudienceHandler).Methods(POST)

	//udpate retargetting
	r.HandleFunc("api/fe/v1/updateretargetting/{id}", adapi.UpdateRetargetHandler).Methods(POST)


	log.Fatal(http.ListenAndServe(":8080", r))
}
