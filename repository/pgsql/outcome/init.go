package outcome

import "database/sql"

type OutcomeRepository struct {
	db *sql.DB
}

func NewOutcomeRepository(db *sql.DB) OutcomeRepository {
	return OutcomeRepository{db: db}
}
