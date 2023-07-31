package usecase

import (
	"context"
	"github.com/personal-finance-vercel/domain/entity"
)

type ExpenseUseCase interface {
	GetExpense(ctx context.Context, userId int) ([]*entity.Expense, error)
	AddExpense(ctx context.Context, req *entity.Expense) error
	FindExpense(ctx context.Context, userId, id int) (*entity.Expense, error)
	PutExpenseById(ctx context.Context, req *entity.Expense, userId, id int) error
}
