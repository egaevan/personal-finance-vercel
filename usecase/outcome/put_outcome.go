package outcome

import (
	"context"
	"errors"
	"github.com/personal-finance-vercel/domain/entity"
)

func (o *outcomeInteractor) PutOutcomeById(ctx context.Context, req *entity.Outcome, userId, id int) error {
	res, err := o.outcomeRepo.GetOutcomeById(ctx, userId, id)
	if err != nil {
		return err
	}

	if res.UserId != userId && res.ID != id {
		return errors.New("not valid user")
	}

	err = o.outcomeRepo.PutOutcomeById(ctx, req, userId, id)
	if err != nil {
		return err
	}
	return nil
}
