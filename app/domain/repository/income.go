package repository

import (
	"context"
	"github.com/personal-finance-vercel/app/domain/entity"
)

type IncomeRepository interface {
	GetIncome(ctx context.Context, userId int) ([]*entity.Income, error)
}
