package income

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewIncomeRepository(db *sql.DB) Repository {
	return Repository{db: db}
}
