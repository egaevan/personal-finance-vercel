package outcome

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/personal-finance-vercel/domain/entity"
	"github.com/personal-finance-vercel/repository/mapper"
	"github.com/personal-finance-vercel/repository/models"
)

func (o *OutcomeRepository) GetOutcomeById(ctx context.Context, userId, id int) (*entity.Outcome, error) {
	query := fmt.Sprintf(`select id, user_id, outcome_information, total_outcome, created_at, updated_at from %s where user_id = $1 AND id = $2`, models.Outcome{}.GetTableName())

	outcome := &models.Outcome{}

	err := o.db.QueryRowContext(ctx, query, userId, id).Scan(
		&outcome.ID,
		&outcome.UserId,
		&outcome.OutcomeInformation,
		&outcome.TotalOutcome,
		&outcome.CreatedDate,
		&outcome.UpdatedDate,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("data not found")
	}

	if err != nil {
		return nil, err
	}

	result := mapper.ToEntityOutcome(outcome)

	return result, nil
}
