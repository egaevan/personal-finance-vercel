package income

import (
	"github.com/personal-finance-vercel/domain/usecase"
	incomeRepo "github.com/personal-finance-vercel/repository/pgsql/income"
	incomeUC "github.com/personal-finance-vercel/usecase/income"
)

type incomeHandler struct {
	incomeUseCase usecase.IncomeUseCase
}

func NewHandler(incomeRepo incomeRepo.IncomeRepository) *incomeHandler {
	incomeUseCase := incomeUC.NewIncomeInteractor(incomeRepo)
	return &incomeHandler{
		incomeUseCase: incomeUseCase,
	}
}