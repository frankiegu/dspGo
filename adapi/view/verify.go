package view

import (
	"fmt"
	"mdsp/common/typedef"
)

const (
	verifyBudgetTemp = "[verifyBudget] %v"
)

func verifyBudgetParam(budget *typedef.Budget) (err error) {
	if budget == nil {
		err = fmt.Errorf(verifyBudgetTemp, "budget is nil")
		return
	}

	if budget.UnlimitedEnable == false {
		if budget.TotalBudget == 0 {
			err = fmt.Errorf(verifyBudgetTemp, "limit enable, total budget can't be 0")
			return
		}

		if budget.DailyBudget == 0 {
			err = fmt.Errorf(verifyBudgetTemp, "limit enable, daily budget can't be 0")
			return
		}

		if budget.TotalBudget < budget.TotalBudget {
			err = fmt.Errorf(verifyBudgetTemp, "total budget can't be less than daily budget")
			return
		}

		// check spend_model
	}

	if budget.FreqCappingEnable == true {
		// check interval

		if budget.FreqCapping == 0 {
			err = fmt.Errorf(verifyBudgetTemp, "freq cap can't be 0")
			return
		}
	}

	if len(budget.DayParting) != 7 {
		err = fmt.Errorf(verifyBudgetTemp, "day parting length is not 7")
		return
	}

	if len(budget.Timezone) == 0 {
		err = fmt.Errorf(verifyBudgetTemp, "timezone can't be empty")
		return
	}

	return
}
