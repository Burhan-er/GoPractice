package util

import "cashbook_tracker/model"

func GroupByCategory(entries []model.Entry) map[string]float64 {
	result :=make(map[string]float64)

	for _, entry := range entries{
		if entry.Type == model.INCOME{
			result[entry.Category] += entry.Amount
		}else if entry.Type == model.EXPENSE{
			result[entry.Category] -= entry.Amount
		}
	}

	return result

}