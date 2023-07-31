package mapper

import (
	"github.com/personal-finance-vercel/domain/entity"
	"github.com/personal-finance-vercel/repository/models"
)

func ToModelExpense(req *entity.Expense) *models.Expense {
	return &models.Expense{
		UserId:             req.UserId,
		TotalExpense:       req.TotalExpense,
		ExpenseInformation: req.ExpenseInformation,
	}
}

func ToEntityExpense(req *models.Expense) *entity.Expense {
	return &entity.Expense{
		ID:                 req.ID,
		UserId:             req.UserId,
		TotalExpense:       req.TotalExpense,
		ExpenseInformation: req.ExpenseInformation,
		CreatedDate:        req.CreatedDate,
		UpdatedDate:        req.UpdatedDate,
	}
}

func ToArrEntityExpense(req []*models.Expense) []*entity.Expense {
	res := make([]*entity.Expense, 0)

	for _, mod := range req {
		singRes := ToEntityExpense(mod)
		res = append(res, singRes)
	}

	return res
}
