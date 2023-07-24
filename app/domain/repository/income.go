package repository

import (
	"context"
	"github.com/personal-finance-vercel/app/domain/entity"
)

type IncomeRepository interface {
	GetIncome(ctx context.Context, userId int) ([]*entity.Income, error)
	InsertIncome(ctx context.Context, ett *entity.Income) error
	GetIncomeById(ctx context.Context, userId, id int) (*entity.Income, error)
	PutIncomeById(ctx context.Context, ett *entity.Income, userId, id int) error
}
