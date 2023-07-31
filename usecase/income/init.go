package income

import (
	"github.com/personal-finance-vercel/repository/pgsql/income"
)

type incomeInteractor struct {
	incomeRepo income.Repository
}

func NewIncomeInteractor(
	incomeRepo income.Repository,
) *incomeInteractor {
	return &incomeInteractor{
		incomeRepo: incomeRepo,
	}
}
