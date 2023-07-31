package income

import (
	"github.com/personal-finance-vercel/domain/usecase"
	expenseRepo "github.com/personal-finance-vercel/repository/pgsql/expense"
	expenseUC "github.com/personal-finance-vercel/usecase/expense"
)

type expenseHandler struct {
	expenseUseCase usecase.ExpenseUseCase
}

func NewHandler(expenseRepo expenseRepo.Repository) *expenseHandler {
	expenseUseCase := expenseUC.NewExpenseInteractor(expenseRepo)
	return &expenseHandler{
		expenseUseCase: expenseUseCase,
	}
}
