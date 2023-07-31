package expense

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/personal-finance-vercel/domain/entity"
	"github.com/personal-finance-vercel/repository/mapper"
	"github.com/personal-finance-vercel/repository/models"
)

func (o *Repository) GetExpenseById(ctx context.Context, userId, id int) (*entity.Expense, error) {
	query := fmt.Sprintf(`select id, user_id, expense_information, total_expense, created_at, updated_at from %s where user_id = $1 AND id = $2`, models.Expense{}.GetTableName())

	expense := &models.Expense{}

	err := o.db.QueryRowContext(ctx, query, userId, id).Scan(
		&expense.ID,
		&expense.UserId,
		&expense.ExpenseInformation,
		&expense.TotalExpense,
		&expense.CreatedDate,
		&expense.UpdatedDate,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("data not found")
	}

	if err != nil {
		return nil, err
	}

	result := mapper.ToEntityExpense(expense)

	return result, nil
}
