package models

import "time"

type Outcome struct {
	ID                 int       `json:"id"`
	UserId             int       `json:"user_id"`
	TotalOutcome       int       `json:"total_outcome"`
	OutcomeInformation string    `json:"outcome_information"`
	CreatedDate        time.Time `json:"created_at"`
	UpdatedDate        time.Time `json:"updated_at"`
}

func (Outcome) GetTableName() string {
	return `outcome`
}
