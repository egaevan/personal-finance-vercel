package outcome

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/personal-finance-vercel/domain/entity"
	"github.com/personal-finance-vercel/repository/mapper"
	"github.com/personal-finance-vercel/repository/models"
	log "github.com/sirupsen/logrus"
)

func (o *OutcomeRepository) GetOutcome(ctx context.Context, userId int) ([]*entity.Outcome, error) {
	query := fmt.Sprintf(`select id, user_id, outcome_information, total_outcome, created_at, updated_at from %s where user_id = $1`, models.Outcome{}.GetTableName())

	rows, err := o.db.QueryContext(ctx, query, userId)

	if err != nil {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("data not found")
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			log.Error(errRow)
		}
	}()

	resRow := make([]*models.Outcome, 0)
	for rows.Next() {
		t := &models.Outcome{}
		err = rows.Scan(
			&t.ID,
			&t.UserId,
			&t.OutcomeInformation,
			&t.TotalOutcome,
			&t.CreatedDate,
			&t.UpdatedDate,
		)

		if err != nil {
			log.Error(err)
			return nil, err
		}

		resRow = append(resRow, t)
	}

	result := mapper.ToArrEntityOutcome(resRow)

	return result, nil
}
