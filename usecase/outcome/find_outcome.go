package outcome

import (
	"context"
	"github.com/personal-finance-vercel/domain/entity"
)

func (o *outcomeInteractor) FindOutcome(ctx context.Context, userId, id int) (*entity.Outcome, error) {
	res, err := o.outcomeRepo.GetOutcomeById(ctx, userId, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
