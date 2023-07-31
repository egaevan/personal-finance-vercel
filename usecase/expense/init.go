package expense

import "github.com/personal-finance-vercel/repository/pgsql/expense"

type expenseInteractor struct {
	expenseRepo expense.Repository
}

func NewExpenseInteractor(
	expenseRepo expense.Repository,
) *expenseInteractor {
	return &expenseInteractor{
		expenseRepo: expenseRepo,
	}
}
