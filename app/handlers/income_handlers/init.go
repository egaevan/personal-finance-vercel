package income_handlers

import (
	"github.com/personal-finance-vercel/app/domain/usecase"
	"github.com/personal-finance-vercel/app/internal/config"
	incomeRepo "github.com/personal-finance-vercel/app/internal/repository/pgsql/income"
	incomeUC "github.com/personal-finance-vercel/app/internal/usecase/income"
)

type incomeHandler struct {
	incomeUseCase usecase.IncomeUseCase
}

func NewHandler(cfg *config.ServerConfig, incomeRepo incomeRepo.IncomeRepository) *incomeHandler {
	incomeUseCase := incomeUC.NewIncomeInteractor(cfg, incomeRepo)
	return &incomeHandler{
		incomeUseCase: incomeUseCase,
	}
}
