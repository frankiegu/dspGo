package adindex

import (
	"log"
	"encoding/json"

	"common/typedef"
)

func CreateCampaign(ca *typedef.Campaign) error {
	return nil
}


func UpdateBudget(caId uint64, budget *typedef.Budget) error {
	return nil
}

func UpdateTarget(caId uint64, target *typedef.Targets) error {
	return nil
}

func UpdateCreative(caId uint64, crvId uint64, crv *typedef.Creative) error {
	return nil
}


func StartCampaign(caId uint64) error {
	return nil
}

func PauseCampaign(caId uint64) error {
	return nil
}

func ActiveCreative(caId uint64, crvIds []uint64) error {
	return nil
}


func InactiveCreative(caId uint64, crvIds []uint64) error {
	return nil
}

func UpdateBlacklist(caId uint64, bl []byte) error {
	return nil
}

func UpdateWhitelist(caId uint64, wl []byte) error {
	return nil
}

func UpdateRetargetting(caId uint64, retarlist []byte) error {
	return nil
}
