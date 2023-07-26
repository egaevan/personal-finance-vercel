package outcome

import (
	"context"
	"fmt"
	"github.com/personal-finance-vercel/domain/entity"
	"github.com/personal-finance-vercel/repository/mapper"
	"github.com/personal-finance-vercel/repository/models"
)

func (o *OutcomeRepository) InsertOutcome(ctx context.Context, ett *entity.Outcome) error {
	req := mapper.ToModelOutcome(ett)
	query := fmt.Sprintf(`INSERT INTO %s(user_id, total_outcome, outcome_information, created_at) VALUES ($1, $2, $3, NOW())`, models.Outcome{}.GetTableName())

	_, err := o.db.ExecContext(ctx, query, req.UserId, req.TotalOutcome, req.OutcomeInformation)

	if err != nil {
		return err
	}

	return nil
}
