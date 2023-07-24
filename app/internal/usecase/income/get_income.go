package income

import (
	"context"
	"github.com/personal-finance-vercel/app/domain/entity"
)

func (i *incomeInteractor) GetAllIncome(ctx context.Context, userId int) ([]*entity.Income, error) {
	res, err := i.incomeRepo.GetIncome(ctx, userId)
	if err != nil {
		return nil, err
	}
	return res, nil
}
