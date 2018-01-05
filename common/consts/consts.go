package consts

/*
* api --> back-end modules rmq
 */
const (
	RMQ_API_EXCHANGE_NAME = "api2be"
	RMQ_API_EXCHANGE_TYPE = "fanout"
	RMQ_API_ROUTINE_NAME  = ""
	RMQ_API_CONTENTTYPE   = "text/plain"
)

/*
* event --> control center module rmq
 */
const (
	RMQ_EVENT_EXCHANGE_NAME = "event2cc"
	RMQ_EVENT_EXCHANGE_TYPE = "fanout"
	RMQ_EVENT_ROUTING_NAME  = ""
	RMQ_EVENT_CONTENTTYPE   = "text/plain"
)

/* * control center --> ad bidder(dsp) moudle rmq message */
const (
	RMQ_CC_EXCHANGE_NAME = "cc2bidder"
	RMQ_CC_EXCHANGE_TYPE = "fanout"
	RMQ_CC_ROUTING_NAME  = ""
	RMQ_CC_CONTENTTYPE   = "text/plain"
)

/*
*cc2bidder message key define
 */
const (
	CC_KEY_ACTIVECAMPAIGN   = "ActiveCampaign"
	CC_KEY_INACTIVECAMPAIGN = "InactiveCampaign"
)

/*
* api2be message key define
 */

const (
	API_KEY_CREATE_CAMPAIGN = "CreateCampaign"

	API_KEY_UPDATE_BASIC			= "UpdateBasic"
	API_KEY_UPDATE_BUDGET			= "UpdateBudget"
	API_KEY_UPDATE_TARGET			= "UpdateTarget"
	API_KEY_UPDATE_POPUP			= "UpdatePopup"
	API_KEY_UPDATE_BANNER			= "UpdateBanner"
	API_KEY_UPDATE_NATIVE			= "UpdateNative"


	API_KEY_START_CAMPAIGN		= "StartCampaign"
	API_KEY_PAUSE_CAMPAIGN		= "PauseCampaign"
	API_KEY_DELETE_CAMPAIGN		= "DeleteCampaign"


	API_KEY_ACTIVE_CREATIVE			= "ActiveCreative"
	API_KEY_INACTIVE_CREATIVE		= "InactiveCreative"
	API_KEY_APPROVE_CREATIVE		= "ApproveCreative"

	API_KEY_UPDATE_WHITEINVENTORY		= "UpdateWhiteInventory"
	API_KEY_UPDATE_BLACKINVENTORY		= "UpdateBlackInventory"

	API_KEY_UPDATE_RETARGETTING_INCLUDE	= "UpdateIncludeRetargetting"
	API_KEY_UPDATE_RETARGETTING_EXCLUDE	= "UpdateExcludeRetargetting"
)

/*
* event2cc message key define
 */



/*
* os consts
*/
/*
const (
	AIX				= "AIX"
	AMIGA	    = "Amiga OS"
	ANDROID		= "Android"
	AROS			=	"AROS"
  BADA      = "Bada"
	BADAOS	  = "Bada OS"
	BEOS      = "BeOS"
BlackBerry OS
Brew
BSD
DangerOS
Firefox OS
Haiku OS
Hiptop OS
HP-UX
IOS
IRIX
JVM
Linux
Linux Smartphone
LiveArea
Mac OS
MeeGo
MorphOS
MTK/Nucleus OS
Nintendo
Nintendo Wii
OpenVMS
OS X
OS/2
Palm OS
Plan 9
PlayStation OS
Rex Qualcomm OS
RIM OS
RIM Tablet OS
RISK OS
Sailfish
SkyOS
Solaris
Syllable
Symbian OS
Tizen
Unknown
WebOS
Windows
Windows CE
Windows Mobile O
Windows Phone OS
Windows RT
Xbox OS
XrossMediaBar (X
)

*/

const (
	OPENRTB_DEV_MOBILE_TABLET	=	1

	OPENRTB_DEV_PC			= 2
	OPENRTB_DEV_TV			= 3
	OPENRTB_DEV_PHONE		= 4
	OPENRTB_DEV_TABLET	= 5

	OPENRTB_DEV_CONNED_DEVICE = 6
	OPENRTB_DEV_SETTOPBOX = 7
)

const (
	OPENRTB_CONNTYPE_UNKNOWN  = 0
	OPENRTB_CONNTYPE_ETHERNET = 1
	OPENRTB_CONNTYPE_WIFI     = 2

	OPENRTB_CONNTYPE_CELLULAR_UNKNOWN = 3

	OPENRTB_CONNTYPE_CELLULAR_2G = 4
	OPENRTB_CONNTYPE_CELLULAR_3G = 5
	OPENRTB_CONNTYPE_CELLULAR_4G = 6
)

const (

)
