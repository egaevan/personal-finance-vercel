package expense

import (
	"context"
	"github.com/personal-finance-vercel/domain/entity"
)

func (o *expenseInteractor) FindExpense(ctx context.Context, userId, id int) (*entity.Expense, error) {
	res, err := o.expenseRepo.GetExpenseById(ctx, userId, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
