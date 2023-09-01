package budget

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewBudgetRepository(db *sql.DB) Repository {
	return Repository{db: db}
}
