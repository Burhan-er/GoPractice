package main

import (
	"cashbook_tracker/model"
	"cashbook_tracker/util"
	"fmt"
)

func main() {
	entries := []model.Entry{
		{Amount: 1000, Type: model.INCOME, Category: "Salary"},
		{Amount: 213, Type: model.INCOME, Category: "Freelance"},
		{Amount: 800, Type: model.INCOME, Category: "Salary"},
		{Amount: 133, Type: model.INCOME, Category: "Freelance"},
		{Amount: 730, Type: model.INCOME, Category: "Salary"},
		{Amount: 644, Type: model.INCOME, Category: "Freelance"},
		{Amount: 750, Type: model.INCOME, Category: "Salary"},
		{Amount: 160, Type: model.EXPENSE, Category: "Transport"},
		{Amount: 1200, Type: model.EXPENSE, Category: "Shopping"},
		{Amount: 550, Type: model.EXPENSE, Category: "Groceries"},
		{Amount: 170, Type: model.EXPENSE, Category: "Transport"},
		{Amount: 1300, Type: model.EXPENSE, Category: "Shopping"},
		{Amount: 600, Type: model.EXPENSE, Category: "Groceries"},
	}

	result := util.GroupByCategory(entries)
	var total float64
	total = 0
	for index, res := range result {
		fmt.Printf("Category: %-10s Total: %0.2f\n", index, res)
		total += res
	}

	fmt.Println("This month's Income: ", total)

}
