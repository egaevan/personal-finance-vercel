package outcome

import (
	"context"
	"github.com/personal-finance-vercel/domain/entity"
)

func (o *outcomeInteractor) GetOutcome(ctx context.Context, userId int) ([]*entity.Outcome, error) {
	res, err := o.outcomeRepo.GetOutcome(ctx, userId)
	if err != nil {
		return nil, err
	}
	return res, nil
}
