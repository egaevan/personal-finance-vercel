package usecase

import (
	"context"
	"github.com/personal-finance-vercel/app/domain/entity"
)

type IncomeUseCase interface {
	GetAllIncome(ctx context.Context, userId int) ([]*entity.Income, error)
}
