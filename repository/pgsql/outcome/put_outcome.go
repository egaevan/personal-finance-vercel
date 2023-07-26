package outcome

import (
	"context"
	"fmt"
	"github.com/personal-finance-vercel/domain/entity"
	"github.com/personal-finance-vercel/repository/mapper"
	"github.com/personal-finance-vercel/repository/models"
)

func (o *OutcomeRepository) PutOutcomeById(ctx context.Context, ett *entity.Outcome, userId, id int) error {
	req := mapper.ToModelOutcome(ett)
	query := fmt.Sprintf(`UPDATE %s SET total_outcome = $1, outcome_information = $2 where user_id = $3 AND id = $4`, models.Outcome{}.GetTableName())

	_, err := o.db.ExecContext(ctx, query, req.TotalOutcome, req.OutcomeInformation, userId, id)

	if err != nil {
		return err
	}

	return nil
}
