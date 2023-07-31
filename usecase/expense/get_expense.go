package expense

import (
	"context"
	"github.com/personal-finance-vercel/domain/entity"
)

func (o *expenseInteractor) GetExpense(ctx context.Context, userId int) ([]*entity.Expense, error) {
	res, err := o.expenseRepo.GetExpense(ctx, userId)
	if err != nil {
		return nil, err
	}
	return res, nil
}
