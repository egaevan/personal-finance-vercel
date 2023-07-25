package income

import (
	"github.com/personal-finance-vercel/repository/pgsql/income"
)

type incomeInteractor struct {
	incomeRepo income.IncomeRepository
}

func NewIncomeInteractor(
	incomeRepo income.IncomeRepository,
) *incomeInteractor {
	return &incomeInteractor{
		incomeRepo: incomeRepo,
	}
}
