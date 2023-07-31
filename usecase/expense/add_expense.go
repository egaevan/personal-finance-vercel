package expense

import (
	"context"
	"github.com/personal-finance-vercel/domain/entity"
)

func (o *expenseInteractor) AddExpense(ctx context.Context, req *entity.Expense) error {
	err := o.expenseRepo.InsertExpense(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
