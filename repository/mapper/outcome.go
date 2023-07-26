package mapper

import (
	"github.com/personal-finance-vercel/domain/entity"
	"github.com/personal-finance-vercel/repository/models"
)

func ToModelOutcome(req *entity.Outcome) *models.Outcome {
	return &models.Outcome{
		UserId:             req.UserId,
		TotalOutcome:       req.TotalOutcome,
		OutcomeInformation: req.OutcomeInformation,
	}
}

func ToEntityOutcome(req *models.Outcome) *entity.Outcome {
	return &entity.Outcome{
		ID:                 req.ID,
		UserId:             req.UserId,
		TotalOutcome:       req.TotalOutcome,
		OutcomeInformation: req.OutcomeInformation,
		CreatedDate:        req.CreatedDate,
		UpdatedDate:        req.UpdatedDate,
	}
}

func ToArrEntityOutcome(req []*models.Outcome) []*entity.Outcome {
	res := make([]*entity.Outcome, 0)

	for _, mod := range req {
		singRes := ToEntityOutcome(mod)
		res = append(res, singRes)
	}

	return res
}
