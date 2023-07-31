package repository

import (
	"context"
	"github.com/personal-finance-vercel/domain/entity"
)

type ExpenseRepository interface {
	GetExpense(ctx context.Context, userId int) ([]*entity.Expense, error)
	InsertExpense(ctx context.Context, ett *entity.Expense) error
	GetExpenseById(ctx context.Context, userId, id int) (*entity.Expense, error)
	PutExpenseById(ctx context.Context, ett *entity.Expense, userId, id int) error
}
