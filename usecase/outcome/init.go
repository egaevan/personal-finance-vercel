package outcome

import (
	"github.com/personal-finance-vercel/repository/pgsql/outcome"
)

type outcomeInteractor struct {
	outcomeRepo outcome.OutcomeRepository
}

func NewOutcomeInteractor(
	outcomeRepo outcome.OutcomeRepository,
) *outcomeInteractor {
	return &outcomeInteractor{
		outcomeRepo: outcomeRepo,
	}
}
