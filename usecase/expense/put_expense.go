package expense

import (
	"context"
	"errors"
	"github.com/personal-finance-vercel/domain/entity"
)

func (o *expenseInteractor) PutExpenseById(ctx context.Context, req *entity.Expense, userId, id int) error {
	res, err := o.expenseRepo.GetExpenseById(ctx, userId, id)
	if err != nil {
		return err
	}

	if res.UserId != userId && res.ID != id {
		return errors.New("not valid user")
	}

	err = o.expenseRepo.PutExpenseById(ctx, req, userId, id)
	if err != nil {
		return err
	}
	return nil
}
