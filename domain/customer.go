package domain

type Customer struct {
	Id          string
	Name        string
	City        string
	ZipCode     string
	Status      string
	DateOfBirth string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
