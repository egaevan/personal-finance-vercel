package usecase

import (
	"context"
	"github.com/personal-finance-vercel/domain/entity"
)

type OutcomeUseCase interface {
	GetOutcome(ctx context.Context, userId int) ([]*entity.Outcome, error)
	AddOutcome(ctx context.Context, req *entity.Outcome) error
	FindOutcome(ctx context.Context, userId, id int) (*entity.Outcome, error)
	PutOutcomeById(ctx context.Context, req *entity.Outcome, userId, id int) error
}
