package repository

import (
	"context"
	"github.com/personal-finance-vercel/domain/entity"
)

type OutcomeRepository interface {
	GetOutcome(ctx context.Context, userId int) ([]*entity.Outcome, error)
	InsertOutcome(ctx context.Context, ett *entity.Outcome) error
	GetOutcomeById(ctx context.Context, userId, id int) (*entity.Outcome, error)
	PutOutcomeById(ctx context.Context, ett *entity.Outcome, userId, id int) error
}
