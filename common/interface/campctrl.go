package common


/*
*
*/
type CampController interface {
	CampStatusController
	CampDevFreqController
	CampBudgetController
	CampAppBudgetController
	CampSiteBudgetController
	CampPeroidControler
}

/*
* campaign status controller
*/
type CampStatusController interface {
	IsActive(camId uint64) (bool, error)

	ActiveCamp(camIds		...uint64)	 error
	InactiveCamp(camIds ...uint64)  error
}

/*
* device freqency of Campaign controller
*/
type CampDevFreqController interface {
	IsDevCamFreqValid(devId string, camId uint64) (bool, error)

	ValidDevFreqCamp(devId string, camId uint64)     error
	InvalidDevFreqInCamp(devId string, camId uint64) error
}

/*
* Campaign buddget controller
*/
type CampBudgetController {
	IsCampBudgetValid(camId uint64)		(bool, error)

	ValidCampBudget(camIds	  ...uint64)	error
	InvalidCampBudget(camIds  ...uint64)	error
}

/*
* campaign placement(app) budget controller
* placement: app bundle id 
*/
type CampAppBudgetController {
	IsCampAppValid(camId uint64, bundleId string)	(bool, error)

	ValidCampAppBudget(camId uint64,   bundleId string)	 error
	InvalidCampAppBudget(camId uint64, bundleId string)  error
}

/*
* campaign placement(site) budget controller
* placement: site domain 
*/
type CampSiteBudgetController {
	IsCampSiteValid(camId uint64, domain string)	(bool, error)

	ValidCampSiteBudget(camId uint64,   domain string)	error
	InvalidCampSiteBudget(camId uint64, domain string)  error
}

/*
* Campaign peroid controller
*/
type CampPeroidControler {
	IsCampPeroidValid(camId uint64)		(bool, error)

	ValidCampInPeroid(camIds   ...uint64)		error
	InvalidCampInPeroid(camIds ...uint64)    error
}
