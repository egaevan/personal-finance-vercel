package outcome

import (
	"context"
	"github.com/personal-finance-vercel/domain/entity"
)

func (o *outcomeInteractor) AddOutcome(ctx context.Context, req *entity.Outcome) error {
	err := o.outcomeRepo.InsertOutcome(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
