package usecase

import (
	"context"
	"github.com/personal-finance-vercel/app/domain/entity"
)

type IncomeUseCase interface {
	GetIncome(ctx context.Context, userId int) ([]*entity.Income, error)
	AddIncome(ctx context.Context, req *entity.Income) error
	FindIncome(ctx context.Context, userId, id int) (*entity.Income, error)
	PutIncomeById(ctx context.Context, req *entity.Income, userId, id int) error
}
