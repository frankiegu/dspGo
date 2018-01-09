package main

import (
	"errors"
	"fmt"
	//libredis "github.com/garyburd/redigo/redis"
	"golang.org/x/net/context"

	"mdsp/utils/redis"
)

const (
	CampKeyTemp                     = "camp_switch_%v"          // <camp_id>, 0:off, 1: on
	CampBudgetLimitSwitchKeyTemp    = "camp_budget_switch_%v"   // <camp_id>, 0: off, 1: on
	DailyBudgetKeyTemp              = "daily_budget_%v"         // <camp_id>
	DailyBudgetSpendStrategyKeyTemp = "daily_spend_strategy_%v" // <camp_id>; 0: no value, 1: ASAP, 2: smooth
	PlacementBudgetKeyTemp          = "placement_cap_%v"        // <camp_id>; can be empty; value: -1 or key not exists: no cap; can update from positive to negative
	TotalBudgetKeyTemp              = "total_budget_%v"         // <camp_id>; key not exists or value -1: no limit

	FreqCapSwitchKeyTemp   = "freq_cap_switch_%v"   // k: <camp_id>; v: 0: off, 1: on
	FreqIntervalCapKeyTemp = "freq_interval_cap_%v" // k:<camp_id>; <interval>: 0: unknown, 1: 3hours, 2: 6hours, 3: 12hours, 4: 24hours; v: <interval>_<cap>, e.g. 2_100
	//FreqCapKeyTemp         = "freq_cap_%v"          // k: <camp_id>; v: freq cap per device
	// freq_statis_<camp_id>_<YYMMDD>_<interval_type>_<interval_no>; interval_type:1 1_0 ... 1_7, interval_type:4 4_0 //

	CreateCampMsgTemp = "[creatCampaignControl] %v"

	// context
	redisHost       = "localhost"
	redisPort       = 6379
	redisContextKey = "redis"
)

var (
	ctx context.Context
)

type CampaignControl struct {
	CampId      uint64 `json:"campId"`
	IsActivePtr *int8  `json:"isActive"` // must, switch
	// MaxBidPtr *uint64 `json:"maxBid"`

	// budget control
	IsBudgetLimitedPtr          *int8   `json:"isBudgetLimited"`          // must, switch
	DailyBudgetPtr              *uint64 `json:"dailyBudget"`              // enable when unlimited budget
	MaxBidPtr                   *uint64 `json:"maxBid"`                   // maxBid <= dailyBudget
	DailyBudgetSpendStrategyPtr *int8   `json:"dailyBudgetSpendStrategy"` // 0: no value, 1: ASAP, 2: smooth
	BudgetPerPlacementPtr       *int64  `json:"budgetPerPlacement"`       // can be nil
	TotalBudgetPtr              *uint64 `json:"totalBudget"`              // not must

	// time control
	FreqCapEnablePtr   *int8   `json:"freqCapEnable"`   // must, 0: disable, 1: enable
	FreqCapIntervalPtr *int8   `json:"freqCapInterval"` // 0: unknown, 1: 3hours, 2: 6hours, 3: 12hours, 4: 24hours
	FreqCapPtr         *uint64 `json:"freqCap"`         // per device

	TimezoneIdPtr     *uint8  `json:"timezoneId"`     // from front end, America/Caracas
	StartTimestampPtr *uint64 `json:"startTimestamp"` // to second
	EndTimestampPtr   *uint64 `json:"endTimestamp"`

	DayParting []uint64 `json:"dayParting"` // Mon - Sun, 0 - 23(24bit): [16777215, 16777215, 16777215, 16777215, 16777215, 16777215, 16777215]
}

func init() {
	ctx = context.Background()

	// redis
	ctx = redis.Open(ctx, redisHost, redisPort, redisContextKey)
}

func getRedisConn() libredis.Conn {
	return redis.GetConn(ctx, redisContextKey)
}

func setCampaignSwitch(camp_id uint64, enable bool) (err error) {
	key := fmt.Sprintf(CampKeyTemp, camp_id)
	conn := getRedisConn()
	if conn == nil {
		return errors.New("setCampaignSwitch redis conn is nil")
	}

	if enable {
		conn.Do("set", key, 1)
	} else {
		conn.Do("set", key, 0)
	}

	return
}

//
func creatCampaignControl(ctrl *CampaignControl) (err error) {
	if ctrl == nil {
		return fmt.Errorf(CreateCampMsgTemp, "Parameter ctrl is nil")
	}

	if ctrl.CampId == 0 {
		return fmt.Errorf(CreateCampMsgTemp, "CampId is missing")
	}

	if ctrl.IsActivePtr == nil {
		return fmt.Errorf(CreateCampMsgTemp, "isActive is missing")
	}

	if ctrl.IsBudgetLimitedPtr == nil {
		return fmt.Errorf(CreateCampMsgTemp, "isBudgetLimited is missing")
	}

	if ctrl.FreqCapEnablePtr == nil {
		return fmt.Errorf(CreateCampMsgTemp, "freqCapEnable is missing")
	}

	conn := getRedisConn()
	if conn == nil {
		return fmt.Errorf(CreateCampMsgTemp, "redis conn is nil")
	}

	// set campaign switch
	if *ctrl.IsActivePtr == 1 {
		err = setCampaignSwitch(ctrl.CampId, true)
	} else {
		err = setCampaignSwitch(ctrl.CampId, false)
	}
	if err != nil {
		return
	}

	budgetSwitchKey := fmt.Sprintf(CampBudgetLimitSwitchKeyTemp, ctrl.CampId)
	// budget control
	if *ctrl.IsBudgetLimitedPtr == 1 {
		if ctrl.DailyBudgetPtr == nil {
			err = fmt.Errorf(CreateCampMsgTemp, "dailyBudget is missing")
			return
		}
		if ctrl.MaxBidPtr == nil {
			err = fmt.Errorf(CreateCampMsgTemp, "maxBid is missing")
			return
		}

		if *ctrl.DailyBudgetPtr < *ctrl.MaxBidPtr {
			err = fmt.Errorf(CreateCampMsgTemp, "dailyBudget can't be less than maxBid")
			return
		}

		// redis set budgetSwitchKey: 1 on

		dailyBudgetKey := fmt.Sprintf(DailyBudgetKeyTemp, ctrl.CampId)
		// redis set daily budget

		spendStrategyKey := fmt.Sprintf(DailyBudgetSpendStrategyKeyTemp, ctrl.CampId)
		// redis set, spend strategy

		// budget per placement
		if ctrl.BudgetPerPlacementPtr != nil {
			placementBudgetKey := fmt.Sprintf(PlacementBudgetKeyTemp, ctrl.CampId)
			// redis set placement budget
		}

		// total budget
		totalBudgetKey := fmt.Sprintf(TotalBudgetKeyTemp, ctrl.CampId)
		// redis set total budget

	} else {
		// redis set budgetSwitchKey: 0 off
	}

	if ctrl.FreqCapEnablePtr != nil && *ctrl.FreqCapEnablePtr == 1 {
		if ctrl.FreqCapPtr == nil {
			err = fmt.Errorf(CreateCampMsgTemp, "frequency cap is missing")
			return
		}

		if ctrl.FreqCapIntervalPtr == nil {
			err = fmt.Errorf(CreateCampMsgTemp, "frequency cap interval is missing")
			return
		}

		// redis: set cap enable

		// redis:

	} else {
		// redis: set cap disable
	}

	return
}

// redis key template

// receive message, distribute to different process

// process impression

// update setup

// spend calc
