package demo

// Get data from the database

type Repository interface {
	GetData() (string, error)
}

type helloRepo struct {
}

func NewHelloRepo() Repository {
	return &helloRepo{}
}

func (r *helloRepo) GetData() (string, error) {
	// Query data from table message
	panic("underconstruction !!")
}
