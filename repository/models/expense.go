package models

import "time"

type Expense struct {
	ID                 int       `json:"id"`
	UserId             int       `json:"user_id"`
	TotalExpense       int       `json:"total_expense"`
	ExpenseInformation string    `json:"expense_information"`
	CreatedDate        time.Time `json:"created_at"`
	UpdatedDate        time.Time `json:"updated_at"`
}

func (Expense) GetTableName() string {
	return `expense`
}
