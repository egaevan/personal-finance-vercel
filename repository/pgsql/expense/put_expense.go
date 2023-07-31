package expense

import (
	"context"
	"fmt"
	"github.com/personal-finance-vercel/domain/entity"
	"github.com/personal-finance-vercel/repository/mapper"
	"github.com/personal-finance-vercel/repository/models"
)

func (o *Repository) PutExpenseById(ctx context.Context, ett *entity.Expense, userId, id int) error {
	req := mapper.ToModelExpense(ett)
	query := fmt.Sprintf(`UPDATE %s SET total_expense = $1, expense_information = $2 where user_id = $3 AND id = $4`, models.Expense{}.GetTableName())

	_, err := o.db.ExecContext(ctx, query, req.TotalExpense, req.ExpenseInformation, userId, id)

	if err != nil {
		return err
	}

	return nil
}
