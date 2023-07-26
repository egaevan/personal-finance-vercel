package income

import (
	"github.com/personal-finance-vercel/domain/usecase"
	outcomeRepo "github.com/personal-finance-vercel/repository/pgsql/outcome"
	outcomeUC "github.com/personal-finance-vercel/usecase/outcome"
)

type outcomeHandler struct {
	outcomeUseCase usecase.OutcomeUseCase
}

func NewHandler(outcomeRepo outcomeRepo.OutcomeRepository) *outcomeHandler {
	outcomeUseCase := outcomeUC.NewOutcomeInteractor(outcomeRepo)
	return &outcomeHandler{
		outcomeUseCase: outcomeUseCase,
	}
}
