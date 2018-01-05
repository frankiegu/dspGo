package adapi

import (
	_ "fmt"
	"strconv"
	"log"
	"strings"
	"net/http"
	"io/ioutil"
	"encoding/json"
)


func CreateCampaignHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("CreateCampaign Request read error"))
		return
	}
	defer r.Body.Close()

	var camp Campaign
	if err := json.Unmarshal(body,&camp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ActiveCreative Request read error"))
		return
	}

	CreateCampaign(&camp)
	log.Println("ids: ", camp)
	w.Write([]byte("ok"))
	w.WriteHeader(http.StatusOK)
}


func UpdateTargetHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.RequestURI, "/")
	campId, _ := strconv.Atoi(path[len(path) -1])

	log.Println("StartCampaignHandler: ", r.RequestURI, campId)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ActiveCreative Request read error"))
		return
	}
	defer r.Body.Close()

	var tar CampTarget
	if err := json.Unmarshal(body,&tar); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ActiveCreative Request read error"))
		return
	}

	UpdateTarget(campId, &tar)
	log.Println("ids: ", tar)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}


func UpdateBasicHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.RequestURI, "/")
	campId, _ := strconv.Atoi(path[len(path) -1])

	log.Println("StartCampaignHandler: ", r.RequestURI, campId)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ActiveCreative Request read error"))
		return
	}
	defer r.Body.Close()

	var basic CampBasic
	if err := json.Unmarshal(body,&basic); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ActiveCreative Request read error"))
		return
	}

	UpdateBasic(campId, &basic)
	log.Println("ids: ", basic)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func UpdateBudgetHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.RequestURI, "/")
	campId, _ := strconv.Atoi(path[len(path) -1])

	log.Println("StartCampaignHandler: ", r.RequestURI, campId)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ActiveCreative Request read error"))
		return
	}
	defer r.Body.Close()

	var budget CampBudget
	if err := json.Unmarshal(body,&budget); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ActiveCreative Request read error"))
		return
	}

	UpdateBudget(campId, &budget)
	log.Println("ids: ", budget)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}


func UpdatePopupHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.RequestURI, "/")
	campId, _ := strconv.Atoi(path[len(path) -1])

	log.Println("StartCampaignHandler: ", r.RequestURI, campId)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ActiveCreative Request read error"))
		return
	}
	defer r.Body.Close()

	type PopupCrvSet struct {
		popupes		[]PopupCrv
	}

	//var popup PopupCrv
	var popupSet PopupCrvSet
	if err := json.Unmarshal(body,&popupSet); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ActiveCreative Request read error"))
		return
	}

	UpdatePopup(campId, popupSet.popupes)
	log.Println("ids: ", popup)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}


func UpdateBannerHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.RequestURI, "/")
	campId, _ := strconv.Atoi(path[len(path) -1])

	log.Println("StartCampaignHandler: ", r.RequestURI, campId)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ActiveCreative Request read error"))
		return
	}
	defer r.Body.Close()

	var banner BannerCrv
	if err := json.Unmarshal(body,&banner); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ActiveCreative Request read error"))
		return
	}

	log.Println("ids: ", banner)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}


func UpdateNativeHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.RequestURI, "/")
	campId, _ := strconv.Atoi(path[len(path) -1])

	log.Println("StartCampaignHandler: ", r.RequestURI, campId)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ActiveCreative Request read error"))
		return
	}
	defer r.Body.Close()

	var native NativeCrv
	if err := json.Unmarshal(body,&native); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ActiveCreative Request read error"))
		return
	}

	log.Println("ids: ", native)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}


func DeleteCampaignHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.RequestURI, "/")
	campId, _ := strconv.Atoi(path[len(path) -1])

	log.Println("StartCampaignHandler: ", r.RequestURI, campId)

	DeleteCampaign(campId)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}


func StartCampaignHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.RequestURI, "/")
	campId, _ := strconv.Atoi(path[len(path) -1])

	log.Println("StartCampaignHandler: ", r.RequestURI, campId)

	StartCampaign(campId)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func PauseCampaignHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.RequestURI, "/")
	campId, _ := strconv.Atoi(path[len(path) -1])

	log.Println("StartCampaignHandler: ", r.RequestURI, campId)

	PauseCampaign(campId)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func ActiveCreativeHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.RequestURI, "/")
	campId, _ := strconv.Atoi(path[len(path) -1])

	log.Println("ActiveCreativeHandler: ", r.RequestURI, campId)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ActiveCreative Request read error"))
		return
	}
	defer r.Body.Close()

	var ids CampCreateiveId

	if err := json.Unmarshal(body,&ids); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ActiveCreative Request read error"))
		return
	}

	log.Println("ids: ", ids)

	ActiveCreative(campId, ids.Ids)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func InactiveCreativeHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.RequestURI, "/")
	campId, _ := strconv.Atoi(path[len(path) -1])

	log.Println("StartCampaignHandler: ", r.RequestURI, campId)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ActiveCreative Request read error"))
		return
	}
	defer r.Body.Close()

	var ids CampCreateiveId

	if err := json.Unmarshal(body,&ids); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ActiveCreative Request read error"))
		return
	}

	log.Println("ids: ", ids)

	InactiveCreative(campId, ids.Ids)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func ApproveCreativeHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.RequestURI, "/")
	campId, _ := strconv.Atoi(path[len(path) -1])

	log.Println("StartCampaignHandler: ", r.RequestURI, campId)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ActiveCreative Request read error"))
		return
	}
	defer r.Body.Close()

	var ids CampCreateiveId

	if err := json.Unmarshal(body,&ids); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ActiveCreative Request read error"))
		return
	}

	log.Println("ids: ", ids)

	ApproveCreative(campId, ids.Ids)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}


func UpdateInventoryHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.RequestURI, "/")
	campId, _ := strconv.Atoi(path[len(path) -1])

	log.Println("StartCampaignHandler: ", r.RequestURI, campId)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ActiveCreative Request read error"))
		return
	}
	defer r.Body.Close()

	var inven Inventory
	if err := json.Unmarshal(body,&inven); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ActiveCreative Request read error"))
		return
	}

	log.Println("inventory: ", inven)

	UpdateInventory(campId, &inven)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func UpdateAudienceHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.RequestURI, "/")
	campId, _ := strconv.Atoi(path[len(path) -1])

	log.Println("StartCampaignHandler: ", r.RequestURI, campId)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ActiveCreative Request read error"))
		return
	}
	defer r.Body.Close()

	var au Audience
	if err := json.Unmarshal(body,&au); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ActiveCreative Request read error"))
		return
	}

	log.Println("audience: ", au)

	UpdateAudience(campId, &au)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func UpdateRetargetHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.RequestURI, "/")
	campId, _ := strconv.Atoi(path[len(path) -1])

	log.Println("StartCampaignHandler: ", r.RequestURI, campId)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ActiveCreative Request read error"))
		return
	}
	defer r.Body.Close()

	var retar CampRetarget
	if err := json.Unmarshal(body,&retar); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ActiveCreative Request read error"))
		return
	}

	log.Println("retargetting: ", retar)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}


func FuncV1Handler(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.RequestURI, "/")
	funcName := path[len(path) - 2]
	campId, _ := strconv.Atoi(path[len(path) -1])

	log.Println("FuncHandler: ", r.RequestURI, funcName, campId)
	switch funcName {
		case "deletecamp":
			DeleteCampaign(campId)
		case "startcamp":
			StartCampaign(campId)
		case "pausecamp":
			PauseCampaign(campId)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ActiveCreative Request read error"))
		return
	}
	defer r.Body.Close()


	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

