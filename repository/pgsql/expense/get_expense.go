package expense

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/personal-finance-vercel/domain/entity"
	"github.com/personal-finance-vercel/repository/mapper"
	"github.com/personal-finance-vercel/repository/models"
	log "github.com/sirupsen/logrus"
)

func (o *Repository) GetExpense(ctx context.Context, userId int) ([]*entity.Expense, error) {
	query := fmt.Sprintf(`select id, user_id, expense_information, total_expense, created_at, updated_at from %s where user_id = $1`, models.Expense{}.GetTableName())

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

	resRow := make([]*models.Expense, 0)
	for rows.Next() {
		t := &models.Expense{}
		err = rows.Scan(
			&t.ID,
			&t.UserId,
			&t.ExpenseInformation,
			&t.TotalExpense,
			&t.CreatedDate,
			&t.UpdatedDate,
		)

		if err != nil {
			log.Error(err)
			return nil, err
		}

		resRow = append(resRow, t)
	}

	result := mapper.ToArrEntityExpense(resRow)

	return result, nil
}
