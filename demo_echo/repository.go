package demo

import "database/sql"

// Get data from the database

type Repository interface {
	GetData() (string, error)
}

type helloRepo struct {
	Db *sql.DB
}

func NewHelloRepo(db *sql.DB) Repository {
	return &helloRepo{
		Db: db,
	}
}

func (r *helloRepo) GetData() (string, error) {
	// Query data from table message
	r.Db.Query("SELECT * FROM message")
	panic("underconstruction !!")
}
