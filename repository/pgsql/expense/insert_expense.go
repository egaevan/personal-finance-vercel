package expense

import (
	"context"
	"fmt"
	"github.com/personal-finance-vercel/domain/entity"
	"github.com/personal-finance-vercel/repository/mapper"
	"github.com/personal-finance-vercel/repository/models"
)

func (o *Repository) InsertExpense(ctx context.Context, ett *entity.Expense) error {
	req := mapper.ToModelExpense(ett)
	query := fmt.Sprintf(`INSERT INTO %s(user_id, total_expense, expense_information, created_at) VALUES ($1, $2, $3, NOW())`, models.Expense{}.GetTableName())

	_, err := o.db.ExecContext(ctx, query, req.UserId, req.TotalExpense, req.ExpenseInformation)

	if err != nil {
		return err
	}

	return nil
}
